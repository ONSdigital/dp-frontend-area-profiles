package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/ONSdigital/dp-api-clients-go/v2/health"
	"github.com/ONSdigital/dp-frontend-area-profiles/config"
	"github.com/ONSdigital/dp-frontend-area-profiles/handlers"
	"github.com/ONSdigital/dp-frontend-area-profiles/routes"
	"github.com/ONSdigital/dp-frontend-area-profiles/utils"
	"github.com/ONSdigital/log.go/v2/log"
	"github.com/gorilla/mux"
)

// Service contains the healthcheck, server and serviceList for the controller
type Service struct {
	Config             *config.Config
	HealthCheck        HealthChecker
	AreaApiHealthCheck *health.Client
	Server             HTTPServer
	ServiceList        *ExternalServiceList
}

// New creates a new service
func New() *Service {
	return &Service{}
}

// Init initialises all the service dependencies, including healthcheck with checkers, api and middleware
func (svc *Service) Init(ctx context.Context, cfg *config.Config, serviceList *ExternalServiceList, clients handlers.Clients, BuildTime, GitCommit, Version string) (err error) {
	log.Info(ctx, "initialising service")

	// Get healthcheck with checkers
	svc.HealthCheck, err = serviceList.GetHealthCheck(cfg, BuildTime, GitCommit, Version)
	if err != nil {
		log.Fatal(ctx, "failed to create health check", err)
		return err
	}
	if err = svc.registerCheckers(ctx, clients); err != nil {
		log.Error(ctx, "failed to register checkers", err)
		return err
	}
	clients.HealthCheckHandler = svc.HealthCheck.Handler

	// Initialise router
	r := mux.NewRouter()
	routes.Setup(ctx, r, cfg, clients)
	if cfg.Debug {
		r.PathPrefix("/areas/dist").Handler(http.StripPrefix("/areas/dist", http.FileServer(http.Dir("assets/dist"))))
	} else {
		r.PathPrefix("/areas/dist").Handler(http.StripPrefix("/areas/dist", http.FileServer(utils.AssetDIR("dist"))))
	}

	svc.Server = serviceList.GetHTTPServer(cfg.BindAddr, r)

	return nil
}

// Run starts an initialised service
func (svc *Service) Run(ctx context.Context, svcErrors chan error) {
	log.Info(ctx, "Starting service", log.Data{"config": svc.Config})

	// Start healthcheck
	svc.HealthCheck.Start(ctx)

	// Start HTTP server
	log.Info(ctx, "Starting server")
	go func() {
		if err := svc.Server.ListenAndServe(); err != nil {
			log.Fatal(ctx, "failed to start http listen and serve", err)
			svcErrors <- err
		}
	}()
}

// Close gracefully shuts the service down in the required order, with timeout
func (svc *Service) Close(ctx context.Context) error {
	log.Info(ctx, "commencing graceful shutdown")
	ctx, cancel := context.WithTimeout(ctx, svc.Config.GracefulShutdownTimeout)
	hasShutdownError := false

	go func() {
		defer cancel()

		// stop healthcheck, as it depends on everything else
		log.Info(ctx, "stop health checkers")
		svc.HealthCheck.Stop()

		// TODO: close any backing services here, e.g. client connections to databases

		// stop any incoming requests
		if err := svc.Server.Shutdown(ctx); err != nil {
			log.Error(ctx, "failed to shutdown http server", err)
			hasShutdownError = true
		}
	}()

	// wait for shutdown success (via cancel) or failure (timeout)
	<-ctx.Done()

	// timeout expired
	if ctx.Err() == context.DeadlineExceeded {
		log.Error(ctx, "shutdown timed out", ctx.Err())
		return ctx.Err()
	}

	// other error
	if hasShutdownError {
		err := errors.New("failed to shutdown gracefully")
		log.Error(ctx, "failed to shutdown gracefully ", err)
		return err
	}

	log.Info(ctx, "graceful shutdown was successful")
	return nil
}

func (svc *Service) registerCheckers(ctx context.Context, c handlers.Clients) (err error) {
	hasErrors := false

	if err = svc.HealthCheck.AddCheck("areas-api", svc.AreaApiHealthCheck.Checker); err != nil {
		hasErrors = true
		log.Error(ctx, "failed to add area-api checker", err)
	}

	if hasErrors {
		return errors.New("Error(s) registering checkers for healthcheck")
	}
	return nil
}

func (svc *Service) IntiateServiceList(config *config.Config, svcList *ExternalServiceList) {
	svc.Config = config
	svc.ServiceList = svcList
	svc.AreaApiHealthCheck = svcList.GetHealthClient("area-api", config.APIRouterURL)
}
