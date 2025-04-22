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

type Config struct {
	HttpPort    string
	TokenExpiry time.Duration
	PrivateKey  *rsa.PrivateKey
	PublicKey   *rsa.PublicKey
}

func LoadConfig() *Config {
	// Environments vars
	log.Printf("Loading environment variables...")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Couldn't load .env file.")
	}

	// Path to auth keys
	var config Config

	// Expiration time
	tokenExpiryStr := os.Getenv("TOKEN_EXPIRY_SECONDS")
	if tokenExpiryStr == "" {
		log.Fatalf("Couldn't find TOKEN_EXPIRY_SECONDS in env")
	}
	seconds, err := strconv.Atoi(tokenExpiryStr)
	if err != nil {
		log.Fatalf("Error while converting TOKEN_EXPIRY_SECONDS to int")
	}
	config.TokenExpiry = time.Duration(seconds) * time.Second

	// Http port
	config.HttpPort = os.Getenv("PORT")
	if config.HttpPort == "" {
		log.Fatalf("Couldn't find PORT in env")
	}

	config.PrivateKey, config.PublicKey = loadKeys(os.Getenv("RS256KEYS_PATH"))
	return &config
}

func loadKeys(keysPath string) (*rsa.PrivateKey, *rsa.PublicKey) {
	privBytes, err := os.ReadFile(keysPath + "private.pem")
	if err != nil {
		log.Fatalf("Error while loading private key. %v", err)
	}

	// var privateKey *rsa.PrivateKey
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privBytes)
	if err != nil {
		log.Fatalf("Error while parsing private .pem data. %v", err)
	}

	pubBytes, err := os.ReadFile(keysPath + "public.pem")
	if err != nil {
		log.Fatalf("Error while loading public key. %v", err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	if err != nil {
		log.Fatalf("Error while parsing public .pem data. %v", err)
	}
	return privateKey, publicKey
}
