// Package main is the entry point for the application.
// Пакет main является точкой входа для приложения.
package main

import (
	"OLO-backend/auth_service/internal/app"
	"OLO-backend/auth_service/internal/config"
	"OLO-backend/pkg/utils/logger"
	"os"
	"os/signal"
	"syscall"

	"log/slog"
)

func main() {
	// Load configuration
	// Загрузка конфигурации
	cfg := config.MustLoad()

	// Setup logger based on environment
	// Настройка логгера в зависимости от окружения
	log := logger.SetupLogger(cfg.Env)

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
