// Package main is the entry point for the OLO-backend OLO-Service application
//
// The api gateway is responsible for routing incoming requests to the appropriate services
// and handling cross-cutting concerns such as authentication, rate limiting, and logging.
package main

import (
	"OLO-backend/olo_service/internal/app"
	"OLO-backend/olo_service/internal/config"
	"OLO-backend/pkg/utils/logger"
	"os"
	"os/signal"
	"syscall"

	"log/slog"
)

func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)

	application := app.New(log, cfg)
	go application.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	osSignal := <-stop
	log.Info("stopping application", slog.String("signal", osSignal.String()))

	application.Stop()
	log.Info("application stopped")
}
