package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func unauthorized(w http.ResponseWriter, msg string) {
	log.Printf("Unauthorized")
	w.Header().Set("content_type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(map[string]string{
		"error":             "invalid_client",
		"error_description": msg,
	})
}
