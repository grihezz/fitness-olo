package app

import (
	"OLO-backend/olo_service/internal/config"
	"OLO-backend/olo_service/internal/service/auth"
	"OLO-backend/olo_service/internal/service/grpc"
	"OLO-backend/olo_service/internal/storage"
	"OLO-backend/olo_service/internal/utils/jwt"
	"log/slog"
)

type App struct {
	GRPCSrv *grpc.Grpc
}

func New(log *slog.Logger, cfg *config.Config) *App {
	var uStorage storage.UserStorage
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

	issuer, err := jwt.NewIssuer("static/private.pem")
	if err != nil {
		panic(err)
	}

	oloService := auth.New(log, uStorage, issuer, cfg.TokenTTL)
	grpcApp := grpc.New(log, cfg.GRPC.Port, oloService)
	return &App{
		GRPCSrv: grpcApp,
	}
}
