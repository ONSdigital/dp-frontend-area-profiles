package handlers

import (
	"context"
	"github.com/ONSdigital/dp-api-clients-go/v2/areas"
	health "github.com/ONSdigital/dp-healthcheck/healthcheck"
	coreModel "github.com/ONSdigital/dp-renderer/model"
	"io"
	"net/http"
)

//go:generate moq -out mocks/areaApi.go . AreaApiClient
//go:generate moq -out mocks/render.go . RenderClient
//go:generate moq -out mocks/renderer.go . RendererClient

type AreaApiClient interface {
	GetArea(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, areaID string) (areaDetails areas.AreaDetails, err error)
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

type RendererClient interface {
	Do(string, []byte) ([]byte, error)
	Checker(ctx context.Context, check *health.CheckState) error
}
