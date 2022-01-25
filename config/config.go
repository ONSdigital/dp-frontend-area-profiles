package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// TODO: remove hello world example config option
// Config represents service configuration for dp-frontend-area-profiles
type Config struct {
	AreaApiURL                 string        `envconfig:"AREA_API_URL"`
	BindAddr                   string        `envconfig:"BIND_ADDR"`
	Debug                      bool          `envconfig:"DEBUG"`
	SiteDomain                 string        `envconfig:"SITE_DOMAIN"`
	PatternLibraryAssetsPath   string        `envconfig:"PATTERN_LIBRARY_ASSETS_PATH"`
	RendererURL                string        `envconfig:"RENDERER_URL"`
	GracefulShutdownTimeout    time.Duration `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT"`
	HealthCheckInterval        time.Duration `envconfig:"HEALTHCHECK_INTERVAL"`
	HealthCheckCriticalTimeout time.Duration `envconfig:"HEALTHCHECK_CRITICAL_TIMEOUT"`
}

var cfg *Config

// Get returns the default config with any modifications through environment
// variables
func Get() (*Config, error) {
	cfg, err := get()
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
	if cfg != nil {
		return cfg, nil
	}

	cfg = &Config{
		BindAddr:                   ":26600",
		Debug:                      false,
		SiteDomain:                 "localhost",
		GracefulShutdownTimeout:    5 * time.Second,
		HealthCheckInterval:        30 * time.Second,
		HealthCheckCriticalTimeout: 90 * time.Second,
	}

	return cfg, envconfig.Process("", cfg)
}
