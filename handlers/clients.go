package handlers

import (
	"context"
	"io"
	"net/http"

	"github.com/ONSdigital/dp-api-clients-go/v2/areas"
	health "github.com/ONSdigital/dp-healthcheck/healthcheck"
	coreModel "github.com/ONSdigital/dp-renderer/model"
)

//go:generate moq -out moq_clients.go -pkg handlers . AreaApiClient RenderClient RendererClient

// AreaApiClient is an interface for requesting area profile specific data
type AreaApiClient interface {
	GetArea(ctx context.Context, userAuthToken string, serviceAuthToken string, collectionID string, areaID string, acceptLang string) (areaDetails areas.AreaDetails, err error)
	GetRelations(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, areaID, acceptLang string) (relations []areas.Relation, err error)
	Checker(ctx context.Context, check *health.CheckState) error
}

// Clients - struct containing all the clients for the controller
type Clients struct {
	HealthCheckHandler func(w http.ResponseWriter, req *http.Request)
	Render             RenderClient
	AreaApi            AreaApiClient
	Renderer           RendererClient
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

// RendererClient is an interface with methods for rending frontend assets
type RendererClient interface {
	Do(string, []byte) ([]byte, error)
	Checker(ctx context.Context, check *health.CheckState) error
}
