package auth

import (
	"backend/database"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	spotify_redirect_uri = "http://localhost:8080/auth/callback"
	access_token         = "" // Global variable to store the access token
)

/*
	LoginHandler | ROUTE: /auth/login

[REQUEST-1] : [Application] -> [Spotify Accounts Service]

	Description: Request authorization to access user data

	GET Request:
	  Endpoint: /authorize
	  QUERY parameters:
	    client_id
	    response_type="code"
	    redirect_uri
	    state (optional)
	    scope

	[Spotify Accounts Service] -> [User]
	  Description: User is prompted to login and authorize access to data by application

	  If user authorizes access, then:
	    User is redirected to *redirect_uri* specified in App setting (Spotify Account Dashboard),
	    returning user back to the application, triggering response.

	[RESPONSE-1] : [Application] <- [User]
	  Description: Response sent from Spotify Accounts Service to Application

	  QUERY parameters:
	    code - authorization code (to be exchnaged for access token)
	    state - value of the state parameter supplied in the request.
*/
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)

	scope := "streaming user-read-email user-read-private"

	authQueryParameters := url.Values{}
	authQueryParameters.Add("response_type", "code")
	authQueryParameters.Add("client_id", os.Getenv("spotify_client_id"))
	authQueryParameters.Add("scope", scope)
	authQueryParameters.Add("redirect_uri", spotify_redirect_uri)

	http.Redirect(w, r, "https://accounts.spotify.com/authorize/?"+authQueryParameters.Encode(), http.StatusSeeOther)
}

/*
CallbackHandler | ROUTE: /auth/callback
*/
func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)

	queryParams := r.URL.Query()
	code := queryParams.Get("code")
	_ = queryParams.Get("state")

	data := url.Values{
		"grant_type":   {"authorization_code"},
		"code":         {code},
		"redirect_uri": {spotify_redirect_uri},
	}

	req, _ := http.NewRequest(
		"POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))

	header_val_auth := "Basic " + base64.StdEncoding.EncodeToString(
		[]byte(os.Getenv("spotify_client_id")+":"+os.Getenv("spotify_client_secret")))

	req.Header = http.Header{
		"Content-Type":  {"application/x-www-form-urlencoded"},
		"Authorization": {header_val_auth},
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyData, _ := io.ReadAll(resp.Body)

		fmt.Println(string(bodyData))

		// Note token_type is always "Bearer" according to docs
		type ResponseData struct {
			AccessToken  string `json:"access_token"`
			TokenType    string `json:"token_type"`
			Scope        string `json:"scope"`
			ExpiresIn    int    `json:"expires_in"`
			RefreshToken string `json:"refresh_token"`
		}

		var responseData ResponseData

		err := json.Unmarshal(bodyData, &responseData)

		if err != nil {
			log.Printf("Error parsing JSON response: %v", err)
		}

		// TODO: store in DB with timestamp
		// note user_id is 0
		err = database.AddSpotifyToken(
			database.DB,
			0, responseData.AccessToken,
			time.Now().Unix()+int64(responseData.ExpiresIn),
			responseData.Scope,
			responseData.RefreshToken)
		if err != nil {
			log.Printf("error within CallbackHandler(): %s\n", err)
		}

		// NOTE: response can either be a redirect or contain data but not both

		// redirect client back to home page
		http.Redirect(w, r, "http://localhost:1420/", http.StatusPermanentRedirect)

		// when the client is redirected back to the home page,
		// the client will make a request to /auth/token which will have a valid access_token

	}
}

/*
TokenHandler | ROUTE: /auth/token?user_id=<int>
*/
func TokenHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)

	// Parse user_id query parameter
	user_id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		// If no access token for the user, send a custom JSON response
		response := map[string]string{
			"message": "Invalid user_id",
		}
		w.Header().Set("Content-Type", "application/json")
		jsonData, _ := json.Marshal(response)
		w.Write(jsonData)
		return
	}

	data, err := database.GetSpotifyToken(database.DB, user_id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// If no access token for the user, send a custom JSON response
			response := map[string]string{
				"message": "No access token available for this user",
			}
			w.Header().Set("Content-Type", "application/json")
			jsonData, _ := json.Marshal(response)
			w.Write(jsonData)
			return
		}

		// For other errors, return a 500 Internal Server Error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Printf("sqlite: %+v\n", data)

	// if access token is expired -> run refresh token flow
	if time.Now().Unix() > data.Expires && data.RefreshToken != "" {
		log.Printf("access token expired, running refresh flow...")
		http.Redirect(w, r, "/auth/token/refresh"+"?user_id="+strconv.Itoa(user_id), http.StatusSeeOther)
		return
	} else { // access token is still fresh

		w.Header().Set("Content-Type", "application/json")
		// Marshal token data into JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Send JSON response
		w.Write(jsonData)
	}
}

/*
	TokenRefreshHandler | ROUTE: /auth/token/refresh?user_id=<int>

[REQUEST-3] : [Application] -> [Spotify Accounts Service]

	Description: request a refresh access_token:

	POST Request:
	  Endpoint: /api/token
	  BODY Parameters: (application/x-www-form-urlencoded)
	    grant_type="refresh_token"
	    refresh_token - refresh token returned from inital authorization code exchange

	  HEADER Parameters:
	    Authorization - Basic <base64 encoded client_id:client_secret>
	    Content-Type - application/x-www-form-urlencoded

	[RESPONSE-3] : [Application] <- [Spotify Accounts Service]
	  Description: response body contains new access_token as JSON data
	  JSON data:
	    access_token - access token for API access
	    token_type - "Bearer"
	    scope - list of scopes granted by user associated with access token
	    expires_in - 3600 (seconds)

	    refresh_token ? Docs say "A new refresh token might be returned too"?
*/
func TokenRefreshHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)

	// Parse the query parameters
	queryParams := r.URL.Query()
	userIDParam := queryParams.Get("user_id")
	// Convert the user_id query parameter to an integer
	user_id, _ := strconv.Atoi(userIDParam)

	DBQuery, err := database.GetSpotifyToken(database.DB, user_id)
	if err != nil {
		log.Printf("no token available for user_id: %d", user_id)
	}

	refresh_token := DBQuery.RefreshToken // fetch refresh token from SQLite

	client := &http.Client{}

	data := url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {refresh_token},
	}

	req, _ := http.NewRequest(
		"POST",
		"https://accounts.spotify.com/api/token",
		strings.NewReader(data.Encode()),
	)

	header_val_auth := "Basic " + base64.StdEncoding.EncodeToString(
		[]byte(os.Getenv("spotify_client_id")+":"+os.Getenv("spotify_client_secret")))

	req.Header = http.Header{
		"Content-Type":  {"application/x-www-form-urlencoded"},
		"Authorization": {header_val_auth},
	}

	resp, _ := client.Do(req)

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		bodyData, _ := io.ReadAll(resp.Body)
		log.Printf("Received JSON data: %s\n", bodyData)

		// Note token_type is always "Bearer" according to docs
		type ResponseData struct {
			access_token  string
			token_type    string
			scope         string
			expires_in    int
			refresh_token string
		}

		var responseData ResponseData

		err := json.Unmarshal(bodyData, &responseData)
		if err != nil {
			log.Printf("Error parsing JSON response: %v", err)
		}

		// update row with new access token, expiry, and refresh token
		err = database.UpdateSpotifyToken(database.DB,
			user_id,
			responseData.access_token,
			time.Now().Unix()+int64(responseData.expires_in),
			responseData.refresh_token)

		if err != nil {
			log.Printf("%s\n", err)
		}

		// return new access_token to client
		DBQuery, _ = database.GetSpotifyToken(database.DB, user_id)
		fmt.Printf("%+v\n", DBQuery)

		w.Header().Set("Content-Type", "application/json")
		// Marshal data into JSON
		jsonData, err := json.Marshal(DBQuery)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Send JSON response
		w.Write(jsonData)
	}
}
