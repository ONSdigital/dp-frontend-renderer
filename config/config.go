package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config structure which hold information for configuring the renderer
type Config struct {
	BindAddr                    string        `envconfig:"BIND_ADDR"`
	Debug                       bool          `envconfig:"DEBUG"`
	SiteDomain                  string        `envconfig:"SITE_DOMAIN"`
	PatternLibraryAssetsPath    string        `envconfig:"PATTERN_LIBRARY_ASSETS_PATH"`
	SupportedLanguages          [2]string     `envconfig:"SUPPORTED_LANGUAGES"`
	EnableCookiesControl        bool          `envconfig:"ENABLE_COOKIES_CONTROL"`
	HealthCheckInterval         time.Duration `envconfig:"HEALTHCHECK_INTERVAL"`
	HealthCheckRecoveryInterval time.Duration `envconfig:"HEALTHCHECK_RECOVERY_INTERVAL"`
}

var cfg *Config

// Get the application and returns the configuration structure
func Get() (*Config, error) {
	cfg, err := get()
	if err != nil {
		return nil, err
	}

	if cfg.Debug {
		cfg.PatternLibraryAssetsPath = "http://localhost:9000/dist"
	} else {
		cfg.PatternLibraryAssetsPath = "//cdn.ons.gov.uk/sixteens/d0f23c6"

	}
	return cfg, nil
}

func get() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg = &Config{
		BindAddr:                    ":20010",
		Debug:                       false,
		SiteDomain:                  "ons.gov.uk",
		SupportedLanguages:          [2]string{"en", "cy"},
		EnableCookiesControl:        false,
		HealthCheckInterval:         10 * time.Second,
		HealthCheckRecoveryInterval: 1 * time.Minute,
	}

	return cfg, envconfig.Process("", cfg)
}
