package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	spotifyClientID     string
	spotifyClientSecret string
	spotifyRedirectURI  = "http://localhost:8080/auth/callback"
	accessToken         = ""
)

func disableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// If it's a preflight request, respond immediately
		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	setenv()
	spotifyClientID = os.Getenv("spotify_client_id")
	spotifyClientSecret = os.Getenv("spotify_client_secret")

	if spotifyClientID == "" || spotifyClientSecret == "" {
		fmt.Println("Spotify client ID and/or client secret not found in environment")
		os.Exit(1)
	}

	http.Handle("/auth/login",
		disableCors(
			http.HandlerFunc(authLoginHandler)))

	http.Handle("/auth/callback",
		disableCors(
			http.HandlerFunc(authCallbackHandler)))

	http.Handle("/auth/token",
		disableCors(
			http.HandlerFunc(authTokenHandler)))

	http.Handle("/auth/token/refresh",
		disableCors(
			http.HandlerFunc(authTokenRefreshHandler)))

	log.Printf("Listening at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
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
func authTokenRefreshHandler(w http.ResponseWriter, r *http.Request) {
	refresh_token := "" // access refresh_token

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
		[]byte(spotifyClientID+":"+spotifyClientSecret))

	req.Header = http.Header{
		"Content-Type":  {"application/x-www-form-urlencoded"},
		"Authorization": {header_val_auth},
	}

	resp, _ := client.Do(req)

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		bodyData, _ := ioutil.ReadAll(resp.Body)

		// print for debug
		fmt.Println(string(bodyData))

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
		accessToken = responseData.access_token

		// TODO: store in DB with timestamp
		// note: when storing to DB struct field might need to be exported
		// i.e. first letter capitalized, see struct tags for more info

		// return new access_token to client
		fmt.Fprintf(w, "{ \"access_token\": \"%s\" }", accessToken)
	}

}

/*
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

// /auth/login
func authLoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)

	log.Printf("%s %s", r.Method, r.URL.Path)

	scope := "streaming user-read-email user-read-private"

	authQueryParameters := url.Values{}
	authQueryParameters.Add("response_type", "code")
	authQueryParameters.Add("client_id", spotifyClientID)
	authQueryParameters.Add("scope", scope)
	authQueryParameters.Add("redirect_uri", spotifyRedirectURI)

	http.Redirect(w, r, "https://accounts.spotify.com/authorize/?"+authQueryParameters.Encode(), http.StatusSeeOther)
}

// Request access token
func authCallbackHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)

	queryParams := r.URL.Query()
	code := queryParams.Get("code")
	_ = queryParams.Get("state")

	data := url.Values{
		"grant_type":   {"authorization_code"},
		"code":         {code},
		"redirect_uri": {spotifyRedirectURI},
	}

	req, _ := http.NewRequest(
		"POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))

	header_val_auth := "Basic " + base64.StdEncoding.EncodeToString(
		[]byte(spotifyClientID+":"+spotifyClientSecret))

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
		bodyData, _ := ioutil.ReadAll(resp.Body)

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
		accessToken = responseData.AccessToken

		// TODO: store in DB with timestamp

		// response can either be a redirect or contain data but not both

		/* return new access_token to client
		w.Header().Set("Content-Type", "application/json")
		w.Write(bodyData) */

		// redirect client back to home page
		http.Redirect(w, r, "http://localhost:1420/", http.StatusPermanentRedirect)

		// when the client is redirected back to the home page,
		// the client will make a request to /auth/token which will have a valid access_token

	}
}

// /auth/token
func authTokenHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)

	fmt.Fprintf(w, "{ \"access_token\": \"%s\" }", accessToken)
}
