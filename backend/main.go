package main

import (
	"backend/api/auth"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	spotifyClientID     string
	spotifyClientSecret string
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
			http.HandlerFunc(auth.LoginHandler)))

	http.Handle("/auth/callback",
		disableCors(
			http.HandlerFunc(auth.CallbackHandler)))

	http.Handle("/auth/token",
		disableCors(
			http.HandlerFunc(auth.TokenHandler)))

	http.Handle("/auth/token/refresh",
		disableCors(
			http.HandlerFunc(auth.TokenRefreshHandler)))

	log.Printf("Listening at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
