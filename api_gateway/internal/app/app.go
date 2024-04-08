package app

import (
	"OLO-backend/api_gateway/internal/entity"
	pauth "OLO-backend/auth_service/generated"
	polo "OLO-backend/olo_service/generated"
	"fmt"
	"log/slog"
	"regexp"

	"OLO-backend/api_gateway/internal/config"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
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

func (app *App) initService(s entity.Socket, fn func(conn *grpc.ClientConn)) *grpc.ClientConn {
	clientConn, err := grpc.Dial(fmt.Sprintf("%s:%d", s.Host, s.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to service: %v", err)
	}
	fn(clientConn)
	return clientConn
}

func (app *App) Start() {
	const op = "httpSrv.Start"
	log := app.log.With(
		slog.String("op", op),
		slog.Int("http_port", app.config.HTTP.Port),
	)
	mux := runtime.NewServeMux()
	authConn := app.initService(entity.Socket{
		Host: app.config.AuthService.Host,
		Port: app.config.AuthService.Port,
	}, func(conn *grpc.ClientConn) {
		err := pauth.RegisterAuthHandler(context.Background(), mux, conn)
		if err != nil {
			log.Info("Failed to register service: %v", err)
		}
	})
	defer authConn.Close()

	oloConn := app.initService(entity.Socket{
		Host: app.config.OloService.Host,
		Port: app.config.OloService.Port,
	}, func(conn *grpc.ClientConn) {
		err := polo.RegisterOLOHandler(context.Background(), mux, conn)
		if err != nil {
			log.Info("Failed to register service: %v", err)
		}
	})
	defer oloConn.Close()

	corsMux := cors(mux)

	httpAddr := app.config.HTTP.ToStr()

	log.Info("API Gateway is listening", slog.String("addr", httpAddr))
	if err := http.ListenAndServe(httpAddr, corsMux); err != nil {
		log.Error("failed to serve: %v", err)
	}
}

func allowedOrigin(origin string) bool {
	if viper.GetString("cors") == "*" {
		return true
	}
	if matched, _ := regexp.MatchString(viper.GetString("cors"), origin); matched {
		return true
	}
	return false
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowedOrigin(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}
