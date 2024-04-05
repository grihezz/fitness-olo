package grpc

import (
	"OLO-backend/olo_service/internal/grpc/ologrpc"
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type Grpc struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(log *slog.Logger, port int, oloService ologrpc.OLO) *Grpc {
	gRPCServer := grpc.NewServer()
	ologrpc.SaveWigdet(gRPCServer, oloService)

	return &Grpc{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *Grpc) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *Grpc) Run() error {
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

func (a *Grpc) Stop() {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op)).Info("stopping gRPC server", slog.Int("port", a.port))
	a.gRPCServer.GracefulStop()
}
