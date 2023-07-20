package main

import (
	"github.com/joho/godotenv"
)

func setenv() {
	godotenv.Load(".env", ".bearer_token", ".client")
	// godotenv.Load(".client")
}
