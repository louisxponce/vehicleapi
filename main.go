package main

import (
	"log"
	"net/http"

	"github.com/louisxponce/vehicleapi/auth"
	"github.com/louisxponce/vehicleapi/internal/api"
	"github.com/louisxponce/vehicleapi/internal/config"
	"github.com/louisxponce/vehicleapi/internal/data"
)

func main() {

	cfg := config.LoadConfig()
	authClients := auth.LoadAuthClients()
	clientStore := auth.NewInMemoryStore(authClients)
	dataAccess := data.NewDataAccess()

	r := api.NewRouter(dataAccess, clientStore, cfg)
	log.Printf("Started listening on port %s", cfg.HttpPort)
	log.Fatal(http.ListenAndServe(":"+cfg.HttpPort, r))
}
