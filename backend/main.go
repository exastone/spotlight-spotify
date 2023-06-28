package main

import (
	spotify_api "backend/api"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	setenv()
	//spotify_api.RefreshAccessToken()
	//spotify_api.GetUserInfoPrivate()
	spotify_api.GetUserInfoPublic()
}
