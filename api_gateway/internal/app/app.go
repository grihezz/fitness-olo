package app

import (
	"OLO-backend/api_gateway/internal/config"
	"OLO-backend/api_gateway/internal/entity"
	pauth "OLO-backend/auth_service/generated"
	polo "OLO-backend/olo_service/generated"
	"context"
	"fmt"
	"log/slog"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

type App struct {
	config *config.Config
	log    *slog.Logger
}

func New(log *slog.Logger) (app *App, err error) {
	app = &App{
		config: config.MustLoad(),
		log:    log,
	}
	return
}

func (app *App) initService(s entity.Socket, fn func(formattedAddr string)) {
	fn(fmt.Sprintf("%s:%d", s.Host, s.Port))
}

func (app *App) Start() {
	const op = "httpSrv.Start"

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	logger := app.log.With(
		slog.String("op", op),
		slog.Int("http_port", app.config.HTTP.Port),
	)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	app.initService(entity.Socket{
		Host: app.config.AuthService.Host,
		Port: app.config.AuthService.Port,
	}, func(formattedAddr string) {
		err := pauth.RegisterAuthHandlerFromEndpoint(ctx, mux, formattedAddr, opts)
		if err != nil {
			logger.Error("can`t register service: %v", err)
		}
	})

	app.initService(entity.Socket{
		Host: app.config.OloService.Host,
		Port: app.config.OloService.Port,
	}, func(formattedAddr string) {
		err := polo.RegisterOLOHandlerFromEndpoint(ctx, mux, formattedAddr, opts)
		if err != nil {
			logger.Error("Failed to register service: %v", err)
		}
	})

	withCors := cors.New(cors.Options{
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler(mux)

	httpAddr := app.config.HTTP.ToStr()

	logger.Info("API Gateway is listening", slog.String("addr", httpAddr))
	if err := http.ListenAndServe(httpAddr, withCors); err != nil {
		logger.Error("failed to serve: %v", err)
	}
}
