package auth

import (
	"crypto/rsa"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/louisxponce/vehicleapi/internal/clients"
)

func TokenHandler(clients map[string]clients.AuthClient, privateKey *rsa.PrivateKey, tokenExpiry time.Duration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Extracting credentials")
		clientId := r.FormValue("client_id")
		clientSecret := r.FormValue("client_secret")
		if clientId == "" && clientSecret == "" {
			username, password, ok := r.BasicAuth()
			if ok {
				clientId = username
				clientSecret = password
			}
		}

		client, ok := clients[clientId]
		if !ok || clientSecret != client.Secret {
			unauthorized(w, "Invalid client credentials")
			return
		}
		log.Printf("Creating token")
		token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
			"sub":   clientId,
			"exp":   time.Now().Add(tokenExpiry).Unix(),
			"iat":   time.Now().Unix(),
			"scope": "basic",
		})

		log.Printf("Signing token")
		tokenString, err := token.SignedString(privateKey)
		if err != nil {
			log.Printf("Failed to sign token. %v", err)
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
}

func unauthorized(w http.ResponseWriter, msg string) {
	log.Printf("Unauthorized")
	w.Header().Set("content_type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(map[string]string{
		"error":             "invalid_client",
		"error_description": msg,
	})
}
