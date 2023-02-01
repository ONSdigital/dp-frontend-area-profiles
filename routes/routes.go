package routes

import (
	"context"

	"github.com/ONSdigital/dp-frontend-area-profiles/config"
	"github.com/ONSdigital/dp-frontend-area-profiles/handlers"
	"github.com/ONSdigital/log.go/v2/log"
	"github.com/gorilla/mux"
)

// Setup registers routes for the service
func Setup(ctx context.Context, r *mux.Router, cfg *config.Config, c handlers.Clients) {
	log.Info(ctx, "adding routes")
	r.StrictSlash(true).Path("/health").HandlerFunc(c.HealthCheckHandler)
	r.StrictSlash(true).Path("/areas").Methods("GET").HandlerFunc(handlers.GeographyStart(ctx, *cfg, c))
	r.StrictSlash(true).Path("/areas/{id}").Methods("GET").HandlerFunc(handlers.GetArea(ctx, *cfg, c))
}
