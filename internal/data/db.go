package data

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "vehicles.db")
	if err != nil {
		log.Fatalf("Could not open the database. %v", err)
	}

	_, err = DB.Exec("PRAGMA journal_mode=WAL;")
	if err != nil {
		log.Fatalf("Failed to enable WAL mode: %v", err)
	}

	log.Printf("Database ready.")
}
