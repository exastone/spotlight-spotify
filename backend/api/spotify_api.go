package spotify_api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// const base_url = "https://accounts.spotify.com"

func RefreshAccessToken() {
	authURL := "https://accounts.spotify.com/api/token"
	authData := url.Values{}
	authData.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", authURL, strings.NewReader(authData.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	authHeader := base64.StdEncoding.EncodeToString([]byte(os.Getenv("client_id") + ":" + os.Getenv("client_secret")))
	req.Header.Set("Authorization", "Basic "+authHeader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Handle response
	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Error parsing response:", err)
		return
	}

	// Access the token from the response
	token, ok := response["access_token"].(string)
	if !ok {
		fmt.Println("Access token not found in response")
		return
	}

	fmt.Println("Access Token:", token)

	env, err := godotenv.Unmarshal("BEARER=\"" + token + "\"")
	err = godotenv.Write(env, "./.bearer_token")
	if err != nil {
		log.Fatal(err)
	}

}

func GetUserInfoPrivate() {

	url := "https://accounts.spotify.com/v1/me"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+os.Getenv("BEARER"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func GetUserInfoPublic() {

	url := "https://api.spotify.com/v1/users/{}"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+os.Getenv("BEARER"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
