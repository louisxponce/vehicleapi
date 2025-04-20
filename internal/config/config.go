package config

import (
	"crypto/rsa"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var KeysPath string
var HttpPort string
var TokenExpiry time.Duration
var PrivateKey *rsa.PrivateKey
var PublicKey *rsa.PublicKey

func LoadEnv() {
	// Environments vars
	log.Printf("Loading environment variables...")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Couldn't load .env file.")
	}

	// Path to auth keys
	KeysPath = os.Getenv("RS256KEYS_PATH")

	// Expiration time
	tokenExpiryStr := os.Getenv("TOKEN_EXPIRY_SECONDS")
	if tokenExpiryStr == "" {
		log.Fatalf("Couldn't find TOKEN_EXPIRY_SECONDS in env")
	}
	seconds, err := strconv.Atoi(tokenExpiryStr)
	if err != nil {
		log.Fatalf("Error while converting TOKEN_EXPIRY_SECONDS to int")
	}
	TokenExpiry = time.Duration(seconds) * time.Second

	// Http port
	HttpPort = os.Getenv("PORT")
	if HttpPort == "" {
		log.Fatalf("Couldn't find PORT in env")
	}
}

func LoadKeys() {
	privBytes, err := os.ReadFile(KeysPath + "private.pem")
	if err != nil {
		log.Fatalf("Error while loading private key. %v", err)
	}

	PrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privBytes)
	if err != nil {
		log.Fatalf("Error while parsing private .pem data. %v", err)
	}

	pubBytes, err := os.ReadFile(KeysPath + "public.pem")
	if err != nil {
		log.Fatalf("Error while loading public key. %v", err)
	}
	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	if err != nil {
		log.Fatalf("Error while parsing public .pem data. %v", err)
	}
}
