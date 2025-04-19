package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func loadEnv() {
	// Environments vars
	log.Printf("Loading environment variables...")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Couldn't load .env file.")
	}

	// // JWT Private Key
	// key := os.Getenv("JWT_SECRET")
	// if key == "" {
	// 	log.Fatalf("Couldn't find JWT_SECRET in env")
	// }
	// jwtKey = []byte(key)

	// Path to auth keys
	keysPath = os.Getenv("RS256KEYS_PATH")

	// Expiration time
	tokenExpiryStr := os.Getenv("TOKEN_EXPIRY_SECONDS")
	if tokenExpiryStr == "" {
		log.Fatalf("Couldn't find TOKEN_EXPIRY_SECONDS in env")
	}
	seconds, err := strconv.Atoi(tokenExpiryStr)
	if err != nil {
		log.Fatalf("Error while converting TOKEN_EXPIRY_SECONDS to int")
	}
	tokenExpiry = time.Duration(seconds) * time.Second

	// Http port
	httpPort = os.Getenv("PORT")
	if httpPort == "" {
		log.Fatalf("Couldn't find PORT in env")
	}
}

func loadKeys() {
	privBytes, err := os.ReadFile(keysPath + "private.pem")
	if err != nil {
		log.Fatalf("Error while loading private key. %v", err)
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privBytes)
	if err != nil {
		log.Fatalf("Error while parsing private .pem data. %v", err)
	}

	pubBytes, err := os.ReadFile(keysPath + "public.pem")
	if err != nil {
		log.Fatalf("Error while loading public key. %v", err)
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	if err != nil {
		log.Fatalf("Error while parsing public .pem data. %v", err)
	}
}
