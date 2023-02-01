package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config represents service configuration for dp-frontend-area-profiles
type Config struct {
	BindAddr                    string         `envconfig:"BIND_ADDR"`
	Debug                       bool           `envconfig:"DEBUG"`
	SiteDomain                  string         `envconfig:"SITE_DOMAIN"`
	PatternLibraryAssetsPath    string         `envconfig:"PATTERN_LIBRARY_ASSETS_PATH"`
	RendererURL                 string         `envconfig:"RENDERER_URL"`
	GracefulShutdownTimeout     time.Duration  `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT"`
	HealthCheckInterval         time.Duration  `envconfig:"HEALTHCHECK_INTERVAL"`
	HealthCheckCriticalTimeout  time.Duration  `envconfig:"HEALTHCHECK_CRITICAL_TIMEOUT"`
	ShowONSMap                  bool           `envconfig:"SHOW_ONS_MAP"`
	APIRouterURL                string         `envconfig:"API_ROUTER_URL"`
	CacheUpdateInterval         *time.Duration `envconfig:"CACHE_UPDATE_INTERVAL"`
	EnableNewNavBar             bool           `envconfig:"ENABLE_NEW_NAVBAR"`
	EnableCensusTopicSubsection bool           `envconfig:"ENABLE_CENSUS_TOPIC_SUBSECTION"`
	CensusTopicID               string         `envconfig:"CENSUS_TOPIC_ID"`
	IsPublishingMode            bool           `envconfig:"IS_PUBLISHING_MODE"`
	SupportedLanguages          []string       `envconfig:"SUPPORTED_LANGUAGES"`
	ServiceAuthToken            string         `envconfig:"SERVICE_AUTH_TOKEN"`
}

var cfg *Config

// Get returns the default config with any modifications through environment
// variables
func Get() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	var err error
	cfg, err = get()
	if err != nil {
		return nil, err
	}

	if cfg.Debug {
		cfg.PatternLibraryAssetsPath = "http://localhost:9002/dist/assets"
	} else {
		cfg.PatternLibraryAssetsPath = "//cdn.ons.gov.uk/dp-design-system/613c855"
	}

	return cfg, nil
}

func get() (*Config, error) {
	cfg = &Config{
		BindAddr:                    ":26600",
		Debug:                       false,
		SiteDomain:                  "localhost",
		RendererURL:                 ":",
		GracefulShutdownTimeout:     5 * time.Second,
		HealthCheckInterval:         30 * time.Second,
		HealthCheckCriticalTimeout:  90 * time.Second,
		APIRouterURL:                "http://localhost:23200/v1",
		EnableNewNavBar:             false,
		EnableCensusTopicSubsection: false,
		CensusTopicID:               "4445",
		IsPublishingMode:            false,
		ServiceAuthToken:            "",
		SupportedLanguages:          []string{"en", "cy"},
	}

	return cfg, envconfig.Process("", cfg)
}
