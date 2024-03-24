package app

import (
	"OLO-backend/api_gateway/internal/entity"
	pauth "OLO-backend/auth_service/generated"
	"fmt"

	"OLO-backend/api_gateway/internal/config"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

type App struct {
	config *config.Config
}

func New() (app *App, err error) {
	app = &App{
		config: config.MustLoad(),
	}
	return
}

func (app *App) initService(s entity.Socket, fn func(conn *grpc.ClientConn)) *grpc.ClientConn {
	clientConn, err := grpc.Dial(fmt.Sprintf("%s:%d", s.Host, s.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to service: %v", err)
	}
	fn(clientConn)
	return clientConn
}

func (app *App) Start() {
	mux := runtime.NewServeMux()
	authConn := app.initService(entity.Socket{
		Host: app.config.AuthService.Host,
		Port: app.config.AuthService.Port,
	}, func(conn *grpc.ClientConn) {
		err := pauth.RegisterAuthHandler(context.Background(), mux, conn)
		if err != nil {
			log.Fatalf("Failed to register service: %v", err)
		}
	})
	defer authConn.Close()

	httpAddr := app.config.HTTP.ToStr()
	log.Printf("API Gateway is listening at %s", httpAddr)
	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
