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

func init() {
	setenv()

	spotifyClientID = os.Getenv("spotify_client_id")
	spotifyClientSecret = os.Getenv("spotify_client_secret")
}
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
	http.Handle("/auth/login", disableCors(http.HandlerFunc(authLoginHandler)))
	http.Handle("/auth/callback", disableCors(http.HandlerFunc(authCallbackHandler)))
	http.Handle("/auth/token", disableCors(http.HandlerFunc(authTokenHandler)))

	log.Printf("Listening at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func authLoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)

	scope := "streaming user-read-email user-read-private"

	authQueryParameters := url.Values{}
	authQueryParameters.Add("response_type", "code")
	authQueryParameters.Add("client_id", spotifyClientID)
	authQueryParameters.Add("scope", scope)
	authQueryParameters.Add("redirect_uri", spotifyRedirectURI)

	http.Redirect(w, r, "https://accounts.spotify.com/authorize/?"+authQueryParameters.Encode(), http.StatusSeeOther)
}

func authCallbackHandler(w http.ResponseWriter, r *http.Request) {
	// log.Printf("%s %s", r.Method, r.URL.Path)

	code := r.URL.Query().Get("code")

	data := url.Values{}
	data.Set("code", code)
	data.Set("redirect_uri", spotifyRedirectURI)
	data.Set("grant_type", "authorization_code")

	req, _ := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(spotifyClientID+":"+spotifyClientSecret)))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return
	}

	bodyString := string(bodyBytes)
	log.Println(bodyString)

	if resp.StatusCode == http.StatusOK {
		var result map[string]interface{}

		err = json.Unmarshal(bodyBytes, &result)
		if err != nil {
			log.Printf("Error parsing JSON response: %v", err)
			return
		}

		if token, ok := result["access_token"].(string); ok {
			accessToken = token
			log.Printf("Access token: %s", accessToken)
		} else {
			log.Println("Access token not found in response")
		}

		http.Redirect(w, r, "http://localhost:1420/", http.StatusSeeOther)
	}
}

func authTokenHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)

	fmt.Fprintf(w, "{ \"access_token\": \"%s\" }", accessToken)
}
