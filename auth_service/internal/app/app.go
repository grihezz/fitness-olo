// Package app provides functionality for initializing the application.
// Пакет app предоставляет функциональность для инициализации приложения.
package app

import (
	"OLO-backend/auth_service/internal/config"
	"OLO-backend/auth_service/internal/service/auth"
	"OLO-backend/auth_service/internal/service/grpc"
	"OLO-backend/auth_service/internal/storage"
	"OLO-backend/pkg/utils/jwt"
	"log/slog"
)

// App represents the application.
// App представляет приложение.
type App struct {
	GRPCSrv *grpc.Grpc
}

// New creates a new instance of the application.
// New создает новый экземпляр приложения.
func New(log *slog.Logger, cfg *config.Config) *App {
	var uStorage storage.UserStorage

	// Initialize user storage based on data provider
	// Инициализация хранилища пользователей на основе провайдера данных
	if cfg.DataProvider == "mysql" {
		uStorage = storage.NewInAuthMysqlStorage(log,
			cfg.MySQLSettings.Address,
			cfg.MySQLSettings.Username,
			cfg.MySQLSettings.Password,
			cfg.MySQLSettings.Database,
			cfg.MySQLSettings.Port)
	} else {
		panic("Not not found provider " + cfg.DataProvider)
	}

	// Initialize JWT issuer
	// Инициализация эмитента JWT
	issuer, err := jwt.NewIssuer("static/private.pem")
	if err != nil {
		panic(err)
	}

	// Initialize authentication service
	// Инициализация сервиса аутентификации
	authService := auth.New(log, uStorage, issuer, cfg.TokenTTL)

	// Initialize gRPC application
	// Инициализация gRPC приложения
	grpcApp := grpc.New(log, cfg.GRPC.Port, authService)

	// Return the application instance
	// Возвращение экземпляра приложения
	return &App{
		GRPCSrv: grpcApp,
	}
}
