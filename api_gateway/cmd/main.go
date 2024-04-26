// Package main is the entry point for the OLO-backend API Gateway application
//
// The api gateway is responsible for routing incoming requests to the appropriate services
// and handling cross-cutting concerns such as authentication, rate limiting, and logging.
package main

import (
	"OLO-backend/api_gateway/internal/app"
	slogpretty "OLO-backend/pkg/utils/logger/handlers"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	log := setupLogger("local")
	a, err := app.New(log)
	if err != nil {
		panic(err)
		return
	}
	a.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	osSignal := <-stop
	log.Info("stopping application", slog.String("signal", osSignal.String()))

	log.Info("application stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)
	return slog.New(handler)
}
