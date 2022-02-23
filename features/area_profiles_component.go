package feature

import (
	"context"
	"net/http"

	"github.com/ONSdigital/dp-api-clients-go/v2/areas"
	"github.com/ONSdigital/dp-api-clients-go/v2/health"
	"github.com/ONSdigital/dp-frontend-area-profiles/assets"
	"github.com/ONSdigital/dp-frontend-area-profiles/config"
	"github.com/ONSdigital/dp-frontend-area-profiles/handlers"
	"github.com/ONSdigital/dp-frontend-area-profiles/service"
	"github.com/ONSdigital/dp-frontend-area-profiles/service/mocks"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	render "github.com/ONSdigital/dp-renderer"
)

type AreaProfileComponent struct {
	Config         *config.Config
	errorChan      chan error
	HTTPServer     *http.Server
	ServiceRunning bool
	svc            *service.Service
}

func NewAreaProfilesComponent() (*AreaProfileComponent, error) {
	ctx := context.Background()
	svcErrors := make(chan error, 1)

	c := &AreaProfileComponent{
		errorChan:      svcErrors,
		HTTPServer:     &http.Server{},
		ServiceRunning: false,
	}

	var err error

	c.Config, err = config.Get()
	if err != nil {
		return nil, err
	}

	initMock := &mocks.InitialiserMock{
		DoGetHealthCheckFunc:  c.DoGetHealthcheckOk,
		DoGetHTTPServerFunc:   c.DoGetHTTPServer,
		DoGetHealthClientFunc: c.DoGetHealthClient,
	}

	svcList := service.NewServiceList(initMock)
	cfg := c.Config
	c.svc = service.New()
	c.svc.IntiateServiceList(cfg, svcList)
	// Hardcode the domain otherwise the CI will pickup the dev / prod domain
	cfg.SiteDomain = "localhost"

	rendererClientMock := &handlers.RendererClientMock{}
	areaClientMock := &handlers.AreaApiClientMock{
		GetAreaFunc: func(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, areaID, acceptLang string) (areas.AreaDetails, error) {
			return areas.AreaDetails{}, nil
		},
	}

	clients := handlers.Clients{
		Render:   render.NewWithDefaultClient(assets.Asset, assets.AssetNames, cfg.PatternLibraryAssetsPath, cfg.SiteDomain),
		Renderer: rendererClientMock,
		AreaApi:  areaClientMock,
	}

	if err = c.svc.Init(ctx, c.Config, svcList, clients, "1", "", ""); err != nil {
		return nil, err
	}

	c.svc.Run(ctx, svcErrors)

	c.ServiceRunning = true
	return c, nil
}

func (c *AreaProfileComponent) DoGetHealthcheckOk(cfg *config.Config, buildTime string, gitCommit string, version string) (service.HealthChecker, error) {
	return &mocks.HealthCheckerMock{
		AddCheckFunc: func(name string, checker healthcheck.Checker) error { return nil },
		StartFunc:    func(ctx context.Context) {},
		StopFunc:     func() {},
	}, nil
}

func (c *AreaProfileComponent) DoGetHTTPServer(bindAddr string, router http.Handler) service.HTTPServer {
	c.HTTPServer.Addr = bindAddr
	c.HTTPServer.Handler = router
	return c.HTTPServer
}

func (c *AreaProfileComponent) DoGetHealthClient(name, url string) *health.Client {
	return nil
}
