// Package auth provides functionality for user authentication.
package auth

import (
	"OLO-backend/auth_service/internal/domain/models"
	"OLO-backend/auth_service/internal/storage"
	"OLO-backend/pkg/model"
	"OLO-backend/pkg/utils/jwt"
	"OLO-backend/pkg/utils/logger/sl"
	"context"
	"errors"
	"fmt"
	jwtgo "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"time"
)

// ErrInvalidCredentials indicates invalid user credentials.
var ErrInvalidCredentials = errors.New("invalid credentials")

// Auth represents an authentication service.
type Auth struct {
	log         *slog.Logger
	userStorage storage.UserStorage
	tokenTTL    time.Duration

	issuer    *jwt.Issuer
	validator *jwt.Validator
}

// New creates a new instance of the authentication service.
func New(log *slog.Logger, userStorage storage.UserStorage, jwtIssuer *jwt.Issuer, jwtValidator *jwt.Validator, tokenTTL time.Duration) *Auth {
	return &Auth{
		userStorage: userStorage,
		log:         log,
		tokenTTL:    tokenTTL,
		issuer:      jwtIssuer,
		validator:   jwtValidator,
	}
}

// Login performs user login and returns a JWT token.
func (a *Auth) Login(email string, password string, appID int) (string, error) {
	const op = "auth.Login"

	log := a.log.With(
		slog.String("op", op),
		slog.String("username", email))

	user, err := a.userStorage.GetUserByEmail(email)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			a.log.Warn("user not found", sl.Err(err))

			return "", fmt.Errorf("%s : %w", op, ErrInvalidCredentials)
		}
		a.log.Error("failed to get user", sl.Err(err))
		return "", fmt.Errorf("%s:%w", op, err)
	}

	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(password)); err != nil {
		a.log.Info("invalid credentials", sl.Err(err))
		return "", fmt.Errorf("%s : %w", op, ErrInvalidCredentials)
	}

	if err != nil {
		return "", fmt.Errorf("%s:%w", op, err)
	}
	log.Info("user logged successfully")

	token, err := a.issuer.NewToken(model.TokenUser{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}, a.tokenTTL)

	if err != nil {
		return "", fmt.Errorf("%s:%w", op, err)
	}
	return token, nil
}

// RegisterNewUser registers a new user.
func (a *Auth) RegisterNewUser(email string, pass string) (int64, error) {
	const op = "auth.RegisterNewUser"

	log := a.log.With(
		slog.String("op", op),
		slog.String("email", email))

	log.Info("register new user")

	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	if err != nil {
		log.Error("failed to generate password hash", sl.Err(err))
		return 0, fmt.Errorf("%s:%w", op, err)
	}

	user, err := a.userStorage.GetUserByEmail(email)
	if user != nil {
		log.Error("user already exists")
		return 0, fmt.Errorf("%s : %w", op, storage.ErrUserExist)
	}

	id, err := a.userStorage.SaveUser(email, passHash)
	if err != nil {
		if errors.Is(err, storage.ErrUserExist) {
			log.Error("user already exists", sl.Err(err))
			return 0, fmt.Errorf("%s : %w", op, storage.ErrUserExist)
		}
	}
	log.Info("user registered")
	return id, nil
}

// GetUserInfo get information user.
func (a *Auth) GetUserInfo(ctx context.Context) (*models.User, error) {
	const op = "auth.GetUserInfo"

	token, err := a.getToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s : %w", op, err)
	}

	payloadUser := a.getPayloadUser(token)

	log := a.log.With(
		slog.String("op", op),
		slog.Int64("user_id", payloadUser.ID),
		slog.String("email", payloadUser.Email))

	log.Debug("get user information")

	user, err := a.userStorage.GetUserById(payloadUser.ID)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			a.log.Warn("user not found", sl.Err(err))
			return nil, fmt.Errorf("%s : %w", op, ErrInvalidCredentials)
		}
		a.log.Error("failed to get user", sl.Err(err))
		return nil, fmt.Errorf("%s:%w", op, err)
	}

	return user, nil
}

// getPayloadUser extracts user information from the JWT token.
func (h *Auth) getPayloadUser(token *jwtgo.Token) model.TokenUser {
	id := int64(token.Claims.(jwtgo.MapClaims)["uid"].(float64))
	email := token.Claims.(jwtgo.MapClaims)["email"].(string)
	role := token.Claims.(jwtgo.MapClaims)["role"].(string)

	return model.TokenUser{
		ID:    id,
		Email: email,
		Role:  role,
	}
}

// getToken retrieves and validates the JWT token from the context.
func (a *Auth) getToken(ctx context.Context) (*jwtgo.Token, error) {
	token, err := a.validator.TokenFromContextMetadata(ctx, "Authorization")
	if err != nil {
		return nil, fmt.Errorf("can't get token")
	}
	return token, nil
}
