package app

import (
	grpcapp "auth/internal/app/grpc"
	"auth/internal/serveses/auth"
	"log/slog"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	stroragePath string,
	tokenTTL time.Duration) *App {

	//storage, err := sqlite.New(stroragePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(log, grpcPort, authService)
	return &App{
		GRPCSrv: grpcApp,
	}
}
