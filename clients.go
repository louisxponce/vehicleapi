package main

import (
	"encoding/json"
	"log"
	"os"
)

type Client struct {
	Secret string `json:"secret"`
}

var clients map[string]Client

// Auth
// Reads the contents of the client file into memory
func loadClientIds() {
	log.Printf("loading client information...")
	file, err := os.Open("clients.json")
	if err != nil {
		log.Fatalf("Falied to load clients. %v", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&clients)
	if err != nil {
		log.Fatalf("Failed to parse client data.")
	}
}
