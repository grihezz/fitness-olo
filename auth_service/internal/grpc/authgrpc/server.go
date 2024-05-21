// Package authgrpc provides gRPC server implementation for authentication service.
package authgrpc

import (
	"OLO-backend/auth_service/generated"
	"OLO-backend/auth_service/internal/domain/models"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// emptyValue represents an empty value for comparison.
const emptyValue = 0

// Auth defines methods for authentication.
type Auth interface {
	Login(email string, password string, appID int) (token string, err error)
	RegisterNewUser(email string, password string) (int64, error)
	GetUserInfo(ctx context.Context) (*models.User, error)
}

// serverAPI implements the generated.AuthServer interface.
type serverAPI struct {
	generated.UnimplementedAuthServer
	auth Auth
}

// Register registers the authentication service with the gRPC server.
func Register(gRPC *grpc.Server, auth Auth) {
	generated.RegisterAuthServer(gRPC, &serverAPI{auth: auth})
}

// Login authenticates a user and returns a token.
func (s *serverAPI) Login(ctx context.Context, req *generated.LoginRequest) (*generated.LoginResponse, error) {
	if req.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}
	if req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}
	if req.GetAppId() == emptyValue {
		return nil, status.Error(codes.InvalidArgument, "app id is required")
	}

	token, err := s.auth.Login(req.GetEmail(), req.GetPassword(), int(req.GetAppId()))

	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &generated.LoginResponse{
		Token: token,
	}, nil
}

// Register registers a new user.
func (s *serverAPI) Register(ctx context.Context, req *generated.RegisterRequest) (*generated.RegisterResponse, error) {
	if err := validateRegister(req); err != nil {
		return nil, err
	}

	userID, err := s.auth.RegisterNewUser(req.GetEmail(), req.GetPassword())

	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &generated.RegisterResponse{
		UserId: userID,
	}, nil
}

// GetUserInfo return information about user
func (s *serverAPI) GetUserInfo(ctx context.Context, _ *generated.GetUserInfoRequest) (*generated.GetUserInfoResponse, error) {
	user, err := s.auth.GetUserInfo(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &generated.GetUserInfoResponse{
		UserId:       user.ID,
		Email:        user.Email,
		Role:         user.Role,
		DateRegister: user.DateRegister,
	}, nil
}

// validateRegister validates the registration request.
func validateRegister(req *generated.RegisterRequest) error {
	if req.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "email is required")
	}
	if req.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}
	return nil
}
