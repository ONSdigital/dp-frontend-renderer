package main

import (
	"context"
	"errors"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/ONSdigital/dp-frontend-renderer/config"
	"github.com/ONSdigital/dp-frontend-renderer/service"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	"github.com/ONSdigital/log.go/log"
	"gopkg.in/yaml.v2"
)

const namespace = "dp-frontend-renderer"

var (
	// BuildTime represents the time in which the service was built
	BuildTime string
	// GitCommit represents the commit (SHA-1) hash of the service that is running
	GitCommit string
	// Version represents the version of the service that is running
	Version string
)

var taxonomyRedirects map[string]string

func init() {
	// TODO: We should be building a breadcrumb from the API but for now
	// we will maintain a yaml file which redirects users to the taxonomy
	// base on the dataset ID
	if taxonomyRedirects == nil {
		taxonomyRedirects = make(map[string]string)
		b, err := ioutil.ReadFile("taxonomy-redirects.yml")
		if err != nil {
			log.Event(nil, "could not open taxonomy redirect file", log.Error(err), log.FATAL)
			os.Exit(1)
		}

		if err := yaml.Unmarshal(b, &taxonomyRedirects); err != nil {
			log.Event(nil, "could not unmarshal taxonomy redirects file", log.Error(err), log.FATAL)
			os.Exit(1)
		}
	}
}

func main() {
	log.Namespace = namespace
	ctx := context.Background()

	if err := run(ctx); err != nil {
		log.Event(ctx, "application unexpectedly failed", log.ERROR, log.Error(err))
		os.Exit(1)
	}

	os.Exit(0)
}

func run(ctx context.Context) error {
	// Make os Signal channels
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Get Config
	cfg, err := config.Get()
	if err != nil {
		log.Event(ctx, "unable to retrieve service configuration", log.FATAL, log.Error(err))
		return err
	}
	log.Event(ctx, "got service configuration", log.Data{"config": cfg}, log.INFO)

	// Create healthcheck with versionInfo
	versionInfo, err := healthcheck.NewVersionInfo(BuildTime, GitCommit, Version)
	if err != nil {
		log.Event(ctx, "failed to create versionInfo for healthcheck", log.FATAL, log.Error(err))
		return err
	}
	hc := healthcheck.New(versionInfo, cfg.HealthCheckCriticalTimeout, cfg.HealthCheckInterval)

	hc.Start(ctx)

	// Run the HTTP service
	httpServer := service.Run(ctx, cfg, taxonomyRedirects, &hc)

	// Wait for OS signal, and do graceful shutdown
	signal := <-signals
	log.Event(ctx, "shutting down after os signal received", log.INFO, log.Data{"signal": signal, "shutdown_timeout": cfg.ShutdownTimeout})
	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)

	// track shutdown gracefully closes app
	var gracefulShutdown bool

	// Shutdown goroutine
	go func() {
		defer cancel()
		var hasShutdownError bool

		log.Event(shutdownCtx, "stopping healthcheck", log.INFO)
		hc.Stop()

		log.Event(shutdownCtx, "stopping http server", log.INFO)
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			log.Event(shutdownCtx, "failed to gracefully shutdown http server", log.FATAL, log.Error(err))
			hasShutdownError = true
		}

		if !hasShutdownError {
			gracefulShutdown = true
		}
	}()

	// wait for Shutdown timeout or success (via cancel)
	<-shutdownCtx.Done()

	if shutdownCtx.Err() == context.DeadlineExceeded {
		log.Event(shutdownCtx, "shutdown timeout", log.ERROR, log.Error(shutdownCtx.Err()))
		return shutdownCtx.Err()
	}

	if !gracefulShutdown {
		err = errors.New("failed to shutdown gracefully")
		log.Event(shutdownCtx, "failed to shutdown gracefully ", log.ERROR, log.Error(err))
		return err
	}

	log.Event(shutdownCtx, "done shutdown gracefully", log.INFO, log.Data{"context": shutdownCtx.Err()})
	return nil
}
