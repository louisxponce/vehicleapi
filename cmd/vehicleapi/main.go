package main

import (
	"log"
	"net/http"

	"github.com/louisxponce/vehicleapi/internal/clients"
	"github.com/louisxponce/vehicleapi/internal/config"
	"github.com/louisxponce/vehicleapi/internal/data"
	"github.com/louisxponce/vehicleapi/internal/router"
)

func main() {

	config := config.LoadConfig()
	authClients := clients.LoadAuthClients()
	clientStore := clients.NewInMemoryStore(authClients)
	dataAccess := data.NewDataAccess()

	r := router.NewRouter(dataAccess, clientStore, config)
	log.Printf("Started listening on port %s", config.HttpPort)
	log.Fatal(http.ListenAndServe(":"+config.HttpPort, r))
}
