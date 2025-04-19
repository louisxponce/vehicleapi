package main

import (
	"crypto/rsa"
	"log"
	"net/http"
	"time"

	"github.com/louisxponce/vehicleapi/middleware"
)

type Vehicle struct {
	Id    string `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

var rawVehicleData []Vehicle

// var jwtKey []byte
var keysPath string
var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey
var tokenExpiry time.Duration
var httpPort string

func main() {

	loadEnv()
	loadKeys()
	loadClientIds()

	// Benchmark how long it takes to read the data into memory
	start := time.Now()
	loadData()
	elapsed := time.Since(start)
	log.Printf("Handling time: %s", elapsed)

	log.Printf("Setting up http server")
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/token", tokenHandler)
	mux.HandleFunc("GET /api/vehicles", middleware.AuthMiddleware(publicKey)(getAll))
	mux.HandleFunc("GET /api/vehicles/{id}", middleware.AuthMiddleware(publicKey)(getSingle))
	log.Printf("Started listening on port %s", httpPort)
	http.ListenAndServe(":"+httpPort, mux)
}
