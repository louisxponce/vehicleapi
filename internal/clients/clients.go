package clients

import (
	"encoding/json"
	"log"
	"os"
)

type AuthClient struct {
	Secret string `json:"secret"`
}

// Auth
// Reads the contents of the client file into memory
func LoadClientIds() map[string]AuthClient {
	log.Printf("loading client information...")
	file, err := os.Open("clients.json")
	if err != nil {
		log.Fatalf("Falied to load clients. %v", err)
	}
	defer file.Close()

	var clients map[string]AuthClient
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&clients)
	if err != nil {
		log.Fatalf("Failed to parse client data.")
	}
	return clients
}
