package main

import (
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
var jwtKey []byte
var tokenExpiry time.Duration
var httpPort string

func extractCredentials(r *http.Request) (string, string) {
	clientId := r.FormValue("client_id")
	clientSecret := r.FormValue("client_secret")

	if clientId == "" && clientSecret == "" {
		username, password, ok := r.BasicAuth()

		if ok {
			clientId = username
			clientSecret = password
		}
	}
	return clientId, clientSecret
}

func main() {

	loadEnv()
	loadClientIds()

	// Benchmark how long it takes to read the data into memory
	start := time.Now()
	loadData()
	elapsed := time.Since(start)
	log.Printf("Handling time: %s", elapsed)

	log.Printf("Setting up http server")
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/token", tokenHandler)
	mux.HandleFunc("GET /api/vehicles", middleware.AuthMiddleware(jwtKey)(getAll))
	mux.HandleFunc("GET /api/vehicles/{id}", middleware.AuthMiddleware(jwtKey)(getSingle))
	log.Printf("Started listening on port %s", httpPort)
	http.ListenAndServe(":"+httpPort, mux)
}
