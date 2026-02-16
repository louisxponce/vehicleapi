package router

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/louisxponce/vehicleapi/internal/api"
	"github.com/louisxponce/vehicleapi/internal/auth"
	"github.com/louisxponce/vehicleapi/internal/clients"
	"github.com/louisxponce/vehicleapi/internal/config"
	"github.com/louisxponce/vehicleapi/internal/data"
	"github.com/louisxponce/vehicleapi/internal/middleware"
)

func NewRouter(dataAccess *data.DataAccess, clientStore *clients.InMemoryStore, cfg *config.Config) http.Handler {

	log.Printf("Setting up server routes and auth middleware")
	r := chi.NewRouter()

	// Handler for the provider part
	r.Post("/api/token", auth.TokenHandler(clientStore, cfg.PrivateKey, cfg.TokenExpiry))
	// r.Get("/myip")

	// Handler for the api part
	handler := api.NewApiHandler(dataAccess)
	r.Route("/api/vehicles", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(cfg.PublicKey))
		r.Get("/", handler.GetAll)
		r.Get("/{id}", handler.GetSingle)
	})
	return r
}
