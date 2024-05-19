package tests

import (
	"OLO-backend/auth_service/generated"
	"OLO-backend/auth_service/internal/config"
	"OLO-backend/auth_service/internal/service/auth"
	"OLO-backend/auth_service/internal/service/grpc"
	"OLO-backend/auth_service/internal/storage"
	"OLO-backend/auth_service/internal/utils/jwt"
	"OLO-backend/pkg/utils/logger"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	googlegrpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"time"
)

type AuthSuite struct {
	suite.Suite
	log *slog.Logger

	storage  storage.UserStorage
	services *auth.Auth

	srv *grpc.Grpc

	accessToken string
}

var (
	targetAddrAuth = "localhost:8080"
	tokenTTL       = "1h"
	portSrv        = 8080
)

func (s *AuthSuite) SetupSuite() {
	s.log = logger.SetupLogger(logger.EnvLocal)

	cfg := config.MustLoad()

	s.log.Info("Server not started yet. Starting server...")
	if cfg.DataProvider != "mysql" {
		s.T().Fatalf("provider %s not found...", cfg.DataProvider)
	}

	s.storage = storage.NewInAuthMysqlStorage(s.log,
		"localhost",
		cfg.MySQLSettings.Username,
		cfg.MySQLSettings.Password,
		cfg.MySQLSettings.Database,
		3306,
	)

	issuer, err := jwt.NewIssuer("static/private.pem")
	if err != nil {
		s.T().Fatalf("jwt Issuer not found key: %v", err)
	}

	duration, _ := time.ParseDuration(tokenTTL)
	s.services = auth.New(s.log, s.storage, issuer, duration)

	s.srv = grpc.New(s.log, portSrv, s.services)
	go s.srv.MustRun()

	s.log.Info("SSO OLO App Integration Tests Started")

	s.initData()
}

func (s *AuthSuite) initData() {
	s.register(userRegister)
	s.accessToken = s.login(userLogin)
}

func (s *AuthSuite) register(requestBody *generated.RegisterRequest) {
	conn, err := googlegrpc.Dial(targetAddrAuth, googlegrpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		s.Fail("Failed to create GRPC request")
		return
	}
	defer conn.Close()

	authClient := generated.NewAuthClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = authClient.Register(ctx, requestBody)
	if !assert.Equal(s.T(), err == nil, true) {
		s.T().FailNow()
	}
}

func (s *AuthSuite) login(requestBody *generated.LoginRequest) string {
	conn, err := googlegrpc.Dial(targetAddrAuth, googlegrpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		s.Fail("Failed to create GRPC request")
		return ""
	}
	defer conn.Close()

	authClient := generated.NewAuthClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req, err := authClient.Login(ctx, requestBody)
	if err != nil {
		s.Fail("Failed call method GRPC request")
		return ""
	}

	return req.GetToken()
}
