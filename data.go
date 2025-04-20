package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func loadData() {
	var err error
	db, err = sql.Open("sqlite3", "vehicles.db")
	if err != nil {
		log.Fatalf("Could not open the database. %v", err)
	}

	_, err = db.Exec("PRAGMA journal_mode=WAL;")
	if err != nil {
		log.Fatalf("Failed to enable WAL mode: %v", err)
	}

	log.Printf("Database ready.")
}
