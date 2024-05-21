// Package app is the entry point for the OLO-backend API Gateway application
//
// The api gateway is responsible for routing incoming requests to the appropriate services
// and handling cross-cutting concerns such as authentication, rate limiting, and logging.
package app

import (
	"OLO-backend/olo_service/generated"
	"OLO-backend/olo_service/internal/config"
	"OLO-backend/olo_service/internal/handler"
	"OLO-backend/olo_service/internal/repository"
	"OLO-backend/olo_service/internal/repository/provider"
	"OLO-backend/olo_service/internal/service"
	"OLO-backend/pkg/utils/jwt"
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type App struct {
	log        *slog.Logger
	handler    *handler.OloHandler
	gRPCServer *grpc.Server
	port       int
}

// New is the entry point for the API Gateway application.
// It initializes the repository and several handlers.
func New(log *slog.Logger, cfg *config.Config) (app *App) {
	dbProvider, err := provider.NewMySQLProvider(
		cfg.MySQLSettings.Address,
		cfg.MySQLSettings.Port,
		cfg.MySQLSettings.Username,
		cfg.MySQLSettings.Password,
		cfg.MySQLSettings.Database)
	if err != nil {
		panic(fmt.Errorf("error connect to database: %v", err))
	}
	repos := repository.NewRepository(dbProvider)

	validator, err := jwt.NewValidator("static/public.pem")
	if err != nil {
		panic(err)
	}

	oloService := service.NewOloService(log, repos)
	oloHandler := handler.NewOloHandler(oloService, validator)
	app = &App{
		log:     log,
		handler: oloHandler,
		port:    cfg.GRPC.Port,
	}
	return
}

func (a *App) Start() {
	a.gRPCServer = grpc.NewServer()
	generated.RegisterOLOServer(a.gRPCServer, a.handler)
	if err := a.run(); err != nil {
		panic(err)
	}
}

func (a *App) run() error {
	const op = "grpcapp.Run"
	log := a.log.With(
		slog.String("op", op),
		slog.Int("grpc_port", a.port),
	)

	log.Info("starting gRPC server")

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("failed to listen %s: %w", op, err)
	}

	log.Info("gRPC server is running", slog.String("addr", l.Addr().String()))
	if err := a.gRPCServer.Serve(l); err != nil {
		log.Error("failed to serve gRPC")
	}
	return nil
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op)).Info("stopping gRPC server", slog.Int("port", a.port))
	a.gRPCServer.GracefulStop()
}
