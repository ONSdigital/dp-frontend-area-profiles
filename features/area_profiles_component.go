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
	cfg.SiteDomain = "localhost"

	relationsdata := []areas.Relation{
		{AreaCode: "E12000001", AreaName: "North East", Href: "/areas/E12000001"},
		{AreaCode: "E12000002", AreaName: "North West", Href: "/areas/E12000002"},
		{AreaCode: "E12000003", AreaName: "Yorkshire and The Humbe", Href: "/areas/E12000003"},
	}

	areaData := areas.AreaDetails{
		Code:        "E92000001",
		Name:        "England",
		DateStarted: "Thu, 01 Jan 2009 00: 00: 00 GMT",
		DateEnd:     "",
		WelshName:   "Lloegr",
		Visible:     true,
		AreaType:    "Country",
		Ancestors: []areas.Ancestor{
			{Name: "England", Level: "", Ancestors: nil, Siblings: nil, Children: nil},
			{Name: "North West", Level: "", Ancestors: nil, Siblings: nil, Children: nil},
			{Name: "Manchester", Level: "", Ancestors: nil, Siblings: nil, Children: nil},
			{Name: "Didsbury East", Level: "", Ancestors: nil, Siblings: nil, Children: nil},
		},
	}

	rendererClientMock := &handlers.RendererClientMock{}
	areaClientMock := &handlers.AreaApiClientMock{
		GetAreaFunc: func(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, areaID, acceptLang string) (areas.AreaDetails, error) {
			return areaData, nil
		},
		GetRelationsFunc: func(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, areaID, acceptLang string) ([]areas.Relation, error) {
			return relationsdata, nil
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
	defer c.svc.Close(ctx)

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
