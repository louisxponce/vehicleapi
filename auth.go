package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func tokenHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Extracting credentials")
	clientId, clientSecret := extractCredentials(r)
	client, ok := clients[clientId]
	if !ok || clientSecret != client.Secret {
		unauthorized(w, "Invalid client credentials")
		return
	}
	log.Printf("Creating token")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   clientId,
		"exp":   time.Now().Add(tokenExpiry).Unix(),
		"iat":   time.Now().Unix(),
		"scope": "basic",
	})

	log.Printf("Signing token")
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Failed to sign token", http.StatusInternalServerError)
		return
	}
	log.Printf("Token successfully signed.")
	json.NewEncoder(w).Encode(map[string]string{
		"access_token": tokenString,
		"token_type":   "bearer",
		"expires_in":   strconv.Itoa(int(tokenExpiry.Seconds())),
	})
}
