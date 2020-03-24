package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Configuration structure which hold information for configuring the renderer
type Config struct {
	BindAddr                 string    `envconfig:"BIND_ADDR"`
	Debug                    bool      `envconfig:"DEBUG"`
	SiteDomain               string    `envconfig:"SITE_DOMAIN"`
	PatternLibraryAssetsPath string    `envconfig:"PATTERN_LIBRARY_ASSETS_PATH"`
	SupportedLanguages       [2]string `envconfig:"SUPPORTED_LANGUAGES"`
	EnableCookiesControl     bool      `envconfig:"ENABLE_COOKIES_CONTROL"`
	EnableJsonLd             bool      `envconfig:"ENABLE_JSON_LD"`
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
		BindAddr:             ":20010",
		Debug:                false,
		SiteDomain:           "ons.gov.uk",
		SupportedLanguages:   [2]string{"en", "cy"},
		EnableCookiesControl: false,
		EnableJsonLd:         false,
	}

	return cfg, envconfig.Process("", cfg)
}
