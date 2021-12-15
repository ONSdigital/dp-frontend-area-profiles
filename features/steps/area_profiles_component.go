package steps

import (
	"context"
	componenttest "github.com/ONSdigital/dp-component-test"
	"github.com/ONSdigital/dp-frontend-area-profiles/config"
	"github.com/ONSdigital/dp-frontend-area-profiles/service"
	"github.com/ONSdigital/dp-frontend-area-profiles/service/mocks"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	"github.com/chromedp/chromedp"
	"log"
	"net/http"
	"time"
)

type Chrome struct {
	execAllocatorCanceller context.CancelFunc
	ctxCanceller           context.CancelFunc
	ctx                    context.Context
}

type AreaProfileComponent struct {
	componenttest.ErrorFeature
	apiFeature              *componenttest.APIFeature
	chrome					Chrome
	Config                  *config.Config
	errorChan               chan error
	HTTPServer              *http.Server
	ServiceRunning          bool
	svc                     *service.Service
	waitTimeOut             time.Duration
}

func NewAreaProfilesComponent() (*AreaProfileComponent, error) {
	ctx := context.Background()
	svcErrors := make(chan error, 1)

	c := &AreaProfileComponent{
		errorChan: svcErrors,
		HTTPServer:     &http.Server{},
		ServiceRunning: false,
		waitTimeOut: 10*time.Second,
	}

	c.SetChromeContext()

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
	c.SetChromeContext()
	return c
}

func (c *AreaProfileComponent) Close() error {
	if c.svc != nil && c.ServiceRunning {
		_ = c.svc.Close(context.Background())
		c.ServiceRunning = false
	}
	c.chrome.ctxCanceller()
	c.chrome.execAllocatorCanceller()
	return nil
}

func (c *AreaProfileComponent) BaseURL() string {
	return "http://" + c.Config.SiteDomain + c.Config.BindAddr
}

func (c *AreaProfileComponent) SetChromeContext() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		// set this to false to be able to watch the browser in action
		chromedp.Flag("headless", true),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	c.chrome.execAllocatorCanceller = cancel

	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	c.chrome.ctxCanceller = cancel

	log.Print("re-starting chrome ...")

	c.chrome.ctx = ctx
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

func (c *AreaProfileComponent) RunWithTimeOut(ctx *context.Context, timeout time.Duration, tasks chromedp.Tasks) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		timeoutContext, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return tasks.Do(timeoutContext)
	}
}
