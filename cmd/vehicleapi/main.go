package main

import (
	"log"
	"net/http"

	"github.com/louisxponce/vehicleapi/internal/api"
	"github.com/louisxponce/vehicleapi/internal/auth"
	"github.com/louisxponce/vehicleapi/internal/clients"
	"github.com/louisxponce/vehicleapi/internal/config"
	"github.com/louisxponce/vehicleapi/internal/data"
	"github.com/louisxponce/vehicleapi/internal/middleware"
)

func main() {

	config.LoadEnv()
	config.LoadKeys()
	authClients := clients.LoadClientIds()
	data.InitDB()

	log.Printf("Setting up http server")
	mux := setupRoutes(authClients)
	// mux := http.NewServeMux()
	// mux.HandleFunc("POST /api/token", auth.TokenHandler(clientMap, config.PrivateKey, config.TokenExpiry))
	// mux.HandleFunc("GET /api/vehicles", middleware.AuthMiddleware(config.PublicKey)(api.GetAll(data.DB)))
	// mux.HandleFunc("GET /api/vehicles/{id}", middleware.AuthMiddleware(config.PublicKey)(api.GetSingle(data.DB)))
	// log.Printf("Started listening on port %s", config.HttpPort)
	log.Fatal(http.ListenAndServe(":"+config.HttpPort, mux))
}

func setupRoutes(authClients map[string]clients.AuthClient) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/token", auth.TokenHandler(authClients, config.PrivateKey, config.TokenExpiry))
	mux.HandleFunc("GET /api/vehicles", middleware.AuthMiddleware(config.PublicKey)(api.GetAll(data.DB)))
	mux.HandleFunc("GET /api/vehicles/{id}", middleware.AuthMiddleware(config.PublicKey)(api.GetSingle(data.DB)))
	log.Printf("Started listening on port %s", config.HttpPort)
	return mux
}
