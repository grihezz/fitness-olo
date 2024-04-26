// Package main is the entry point for the application.
// Пакет main является точкой входа для приложения.
package main

import (
	"OLO-backend/auth_service/internal/app"
	"OLO-backend/auth_service/internal/config"
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
	// Load configuration
	// Загрузка конфигурации
	cfg := config.MustLoad()

	// Setup logger based on environment
	// Настройка логгера в зависимости от окружения
	log := setupLogger(cfg.Env)

	// Initialize application
	// Инициализация приложения
	application := app.New(log, cfg)

	// Start gRPC server in a separate goroutine
	// Запуск gRPC сервера в отдельной горутине
	go application.GRPCSrv.MustRun()

	// Listen for OS signals to gracefully shut down the application
	// Ожидание сигналов ОС для грациозной остановки приложения
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	osSignal := <-stop
	log.Info("stopping application", slog.String("signal", osSignal.String()))

	// Stop gRPC server
	// Остановка gRPC сервера
	application.GRPCSrv.Stop()
	log.Info("application stopped")
}

// setupLogger sets up the logger based on the environment.
// Настройка логгера в зависимости от окружения.
func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		// For local environment, use pretty logging
		// Для локальной среды использовать красивое логирование
		log = setupPrettySlog()
	case envDev:
		// For development environment, log debug level messages
		// Для среды разработки логировать сообщения на уровне отладки
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		// For production environment, log info level messages
		// Для продакшен среды логировать сообщения на уровне информации
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}

// setupPrettySlog sets up a pretty logger for local environment.
// Настройка красивого логгера для локальной среды.
func setupPrettySlog() *slog.Logger {
	// Configure pretty logger options
	// Настройка параметров красивого логгера
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	// Create a new pretty logger handler
	// Создание нового обработчика красивого логгера
	handler := opts.NewPrettyHandler(os.Stdout)

	// Return a new logger with the pretty handler
	// Возвращение нового логгера с красивым обработчиком
	return slog.New(handler)
}
