package steps

import (
	"context"
	componenttest "github.com/ONSdigital/dp-component-test"
	"github.com/ONSdigital/dp-frontend-area-profiles/config"
	"github.com/ONSdigital/dp-frontend-area-profiles/service"
	"github.com/ONSdigital/dp-frontend-area-profiles/service/mocks"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	"net/http"
)

type AreaProfileComponent struct {
	componenttest.ErrorFeature
	apiFeature              *componenttest.APIFeature
	Config                  *config.Config
	errorChan               chan error
	HTTPServer              *http.Server
	ServiceRunning          bool
	svc                     *service.Service
}

func NewAreaProfilesComponent() (*AreaProfileComponent, error) {
	ctx := context.Background()
	svcErrors := make(chan error, 1)

	c := &AreaProfileComponent{
		errorChan: svcErrors,
		HTTPServer:     &http.Server{},
		ServiceRunning: false,
	}

	var err error

	c.Config, err = config.Get()
	if err != nil {
		return nil, err
	}

	initMock := &mocks.InitialiserMock{
		DoGetHealthCheckFunc:             c.DoGetHealthcheckOk,
		DoGetHTTPServerFunc:              c.DoGetHTTPServer,
	}

	svcList := service.NewServiceList(initMock)

	c.svc = service.New()
	if err = c.svc.Init(ctx, c.Config, svcList, "1", "", ""); err != nil {
		return nil, err
	}
	c.svc.Run(ctx, svcErrors)

	c.ServiceRunning = true
	c.apiFeature = componenttest.NewAPIFeature(c.InitialiseService)

	return c, nil
}

func (c *AreaProfileComponent) InitialiseService() (http.Handler, error) {
	return c.HTTPServer.Handler, nil
}

func (c *AreaProfileComponent) Reset() *AreaProfileComponent {
	c.apiFeature.Reset()
	return c
}

func (c *AreaProfileComponent) Close() error {
	if c.svc != nil && c.ServiceRunning {
		_ = c.svc.Close(context.Background())
		c.ServiceRunning = false
	}
	return nil
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
