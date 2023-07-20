package auth

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

var spotify_client_id = os.Getenv("spotify_client_id")
var spotify_client_secret = os.Getenv("spotify_client_secret")
var access_token string // Global variable to store the access token

const spotify_redirect_uri = "http://localhost:8080/auth/callback"

func AuthLoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)

	scope := "streaming user-read-email user-read-private"

	authQueryParameters := url.Values{}
	authQueryParameters.Add("response_type", "code")
	authQueryParameters.Add("client_id", spotify_client_id)
	authQueryParameters.Add("scope", scope)
	authQueryParameters.Add("redirect_uri", spotify_redirect_uri)

	http.Redirect(w, r, "https://accounts.spotify.com/authorize/?"+authQueryParameters.Encode(), http.StatusSeeOther)
}

func AuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	// log.Printf("%s %s", r.Method, r.URL.Path)

	code := r.URL.Query().Get("code")

	data := url.Values{}
	data.Set("code", code)
	data.Set("redirect_uri", spotify_redirect_uri)
	data.Set("grant_type", "authorization_code")

	req, _ := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(spotify_client_id+":"+spotify_client_secret)))
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
			access_token = token
			log.Printf("Access token: %s", access_token)
		} else {
			log.Println("Access token not found in response")
		}

		http.Redirect(w, r, "http://localhost:1420/", http.StatusSeeOther)
	}
}

func AuthTokenHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)

	fmt.Fprintf(w, "{ \"access_token\": \"%s\" }", access_token)
}
