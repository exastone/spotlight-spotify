package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitializeDB() (*sql.DB, error) {
	// Initialize connection to the database
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
