package handlers

import (
	"context"
	clients "github.com/ONSdigital/dp-api-clients-go/v2/areas"
	render "github.com/ONSdigital/dp-renderer"
	"io"
	"net/http"

	coreModel "github.com/ONSdigital/dp-renderer/model"
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

// ClientError is an interface that can be used to retrieve the status code if a client has errored
type ClientError interface {
	Error() string
	Code() int
}

// RenderClient is an interface with methods for require for rendering a template
type RenderClient interface {
	BuildPage(w io.Writer, pageModel interface{}, templateName string)
	NewBasePageModel() coreModel.Page
}
