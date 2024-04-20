// Package main is the entry point for the OLO-backend OLO-Service application
//
// The api gateway is responsible for routing incoming requests to the appropriate services
// and handling cross-cutting concerns such as authentication, rate limiting, and logging.
package main

import (
	"OLO-backend/olo_service/internal/app"
	"OLO-backend/olo_service/internal/config"
	"OLO-backend/pkg/utils/logger/handlers"
	"os"
	"os/signal"
	"syscall"

	"log/slog"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)

	application := app.New(log, cfg)
	go application.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	osSignal := <-stop
	log.Info("stopping application", slog.String("signal", osSignal.String()))

	application.Stop()
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
