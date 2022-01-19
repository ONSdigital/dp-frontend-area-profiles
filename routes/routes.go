package routes

import (
	"context"
	"net/http"

	clients "github.com/ONSdigital/dp-api-clients-go/v2/areas"
	"github.com/ONSdigital/dp-frontend-area-profiles/config"
	"github.com/ONSdigital/dp-frontend-area-profiles/handlers"
	render "github.com/ONSdigital/dp-renderer"

	"github.com/ONSdigital/log.go/v2/log"
	"github.com/gorilla/mux"
)

type AreaApiClient interface {
	GetArea(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, areaID string) (areaDetails clients.AreaDetails, err error)
}

// Clients - struct containing all the clients for the controller
type Clients struct {
	HealthCheckHandler func(w http.ResponseWriter, req *http.Request)
	Render             *render.Render
	AreaApi            AreaApiClient
}

// Setup registers routes for the service
func Setup(ctx context.Context, r *mux.Router, cfg *config.Config, c Clients) {
	log.Info(ctx, "adding routes")
	r.StrictSlash(true).Path("/health").HandlerFunc(c.HealthCheckHandler)
	r.StrictSlash(true).Path("/areas").Methods("GET").HandlerFunc(handlers.GeographyStart(*cfg, c.Render))
	r.StrictSlash(true).Path("/areas/{id}").Methods("GET").HandlerFunc(handlers.GetArea(ctx, *cfg, c))
}
