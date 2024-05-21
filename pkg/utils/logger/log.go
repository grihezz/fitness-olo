package logger

import (
	slogpretty "OLO-backend/pkg/utils/logger/handlers"
	"log/slog"
	"os"
)

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

// SetupLogger sets up the logger based on the environment.
// Настройка логгера в зависимости от окружения.
func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case EnvLocal:
		// For local environment, use pretty logging
		// Для локальной среды использовать красивое логирование
		log = setupPrettySlog()
	case EnvDev:
		// For development environment, log debug level messages
		// Для среды разработки логировать сообщения на уровне отладки
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case EnvProd:
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
