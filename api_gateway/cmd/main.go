// Package main is the entry point for the OLO-backend API Gateway application
//
// The api gateway is responsible for routing incoming requests to the appropriate services
// and handling cross-cutting concerns such as authentication, rate limiting, and logging.
package main

import (
	"OLO-backend/api_gateway/internal/app"
	"OLO-backend/pkg/utils/logger"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log := logger.SetupLogger(logger.EnvLocal)
	a, err := app.New(log)
	if err != nil {
		panic(err)
	}
	a.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	osSignal := <-stop
	log.Info("stopping application", slog.String("signal", osSignal.String()))

	log.Info("application stopped")
}
