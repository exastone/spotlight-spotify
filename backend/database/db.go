package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Initialize database connection
func InitializeDB() (*sql.DB, error) {
	var err error
	DB, err = sql.Open("sqlite3", "./storage/test.db")
	if err != nil {
		panic(err)
	}
	// Execute PRAGMA statement to set the journal mode
	_, err = DB.Exec("PRAGMA journal_mode = WAL")
	if err != nil {
		panic(err)
	}
	// Execute PRAGMA statement to set busy timeout duration
	_, err = DB.Exec("PRAGMA busy_timeout = 5000")
	if err != nil {
		panic(err)
	}
	return DB, err
}

func createTokenTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS spotify_tokens (
		user_id INTEGER PRIMARY KEY,
		access_token TEXT,
		expires INTEGER,
		scope TEXT,
		refresh_token TEXT
	);`

	_, err := db.Exec(query)
	return err
}

func AddSpotifyToken(db *sql.DB, user_id int, accessToken string, expires int64, scope string, refreshToken string) error {
	query := `
	INSERT INTO spotify_tokens (user_id, access_token, expires, scope, refresh_token)
	VALUES (?, ?, ?, ?, ?);`

	_, err := db.Exec(query, user_id, accessToken, expires, scope, refreshToken)
	return err
}

type SpotifyToken struct {
	UserID       int    `json:"user_id"`
	AccessToken  string `json:"access_token"`
	Expires      int64  `json:"expires"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
}

func GetSpotifyToken(db *sql.DB, user_id int) (SpotifyToken, error) {
	var token SpotifyToken

	query := `SELECT user_id, access_token, expires, scope, refresh_token
	          FROM spotify_tokens
	          WHERE user_id = ?`

	row := db.QueryRow(query, user_id)
	err := row.Scan(&token.UserID, &token.AccessToken, &token.Expires, &token.Scope, &token.RefreshToken)
	if err != nil {
		return SpotifyToken{}, err
	}
	return token, nil
}

// May want to add scope update as well
func UpdateSpotifyToken(db *sql.DB, user_id int, newAccessToken string, newExpiry int64, newRefreshToken string) error {
	query := `
	UPDATE spotify_tokens
	SET access_token = ?, expires = ?, refresh_token = ?
	WHERE user_id = ?`

	_, err := db.Exec(query, newAccessToken, newExpiry, newRefreshToken, user_id)
	return err
}
