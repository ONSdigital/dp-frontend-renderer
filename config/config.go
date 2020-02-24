package config

import (
	"github.com/ONSdigital/go-ns/log"
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
}

var cfg *Config

// Get the application and returns the configuration structure
func Get() (*Config, error) {
	cfg, err := get()
	if err != nil {
		log.Error(err, nil)
		return nil, err
	}

	if cfg.Debug {
		cfg.PatternLibraryAssetsPath = "http://localhost:9000/dist"
	} else {
		cfg.PatternLibraryAssetsPath = "//cdn.ons.gov.uk/sixteens/f816ac8"

	}
	return cfg, nil
}

func get() (*Config, error) {
	if cfg != nil {
		log.Debug("cfg already set and is not nil", nil)
		return cfg, nil
	}

	cfg = &Config{
		BindAddr:                 ":20010",
		Debug:                    false,
		SiteDomain:               "ons.gov.uk",
		SupportedLanguages:       [2]string{"en", "cy"},
		EnableCookiesControl:     false,
	}

	return cfg, envconfig.Process("", cfg)
}

func (cfg *Config) Log() {
	log.Debug("Configuration", log.Data{
		"BindAddr":                 cfg.BindAddr,
		"Debug":                    cfg.Debug,
		"SiteDomain":               cfg.SiteDomain,
		"PatternLibraryAssetsPath": cfg.PatternLibraryAssetsPath,
		"SupportedLanguages":       cfg.SupportedLanguages,
		"EnableCookiesControl":     cfg.EnableCookiesControl,
	})
}
