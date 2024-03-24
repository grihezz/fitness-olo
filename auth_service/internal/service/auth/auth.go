package auth

import (
	"OLO-backend/auth_service/internal/storage"
	"OLO-backend/auth_service/internal/utils/jwt"
	"OLO-backend/auth_service/internal/utils/logger/sl"
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"time"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type Auth struct {
	log         *slog.Logger
	userStorage storage.UserStorage
	tokenTTL    time.Duration
	issuer      *jwt.Issuer
}

func New(log *slog.Logger, userStorage storage.UserStorage, jwtIssuer *jwt.Issuer, tokenTTL time.Duration) *Auth {
	return &Auth{
		userStorage: userStorage,
		log:         log,
		tokenTTL:    tokenTTL,
		issuer:      jwtIssuer,
	}
}

func (a *Auth) Login(ctx context.Context, email string, password string, appID int) (string, error) {
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

	token, err := a.issuer.NewToken(user, a.tokenTTL)
	if err != nil {
		return "", fmt.Errorf("%s:%w", op, err)
	}
	return token, nil
}

func (a *Auth) RegisterNewUser(ctx context.Context, email string, pass string) (int64, error) {
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

func (a *Auth) IsAdmin(ctx context.Context, userId int64) (bool, error) {
	const op = "auth.IsAdmin"

	log := a.log.With(
		slog.String("op", op),
		slog.Int64("user_id", userId))

	log.Info("check if user is admin")

	user, err := a.userStorage.GetUserById(userId)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			a.log.Warn("user not found", sl.Err(err))

			return false, fmt.Errorf("%s : %w", op, ErrInvalidCredentials)
		}
		a.log.Error("failed to get user", sl.Err(err))
		return false, fmt.Errorf("%s:%w", op, err)
	}

	role := user.Role
	isAdmin := role == "ADMIN"
	return isAdmin, nil
}
