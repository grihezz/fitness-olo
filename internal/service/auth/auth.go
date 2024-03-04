package auth

import (
	"auth/internal/domain/models"
	"auth/internal/lib/jwt"
	"auth/internal/lib/logger/sl"
	"auth/internal/storage"
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"time"
)

var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrInvalidAppID = errors.New("invalid app")

type Auth struct {
	log         *slog.Logger
	usrSaver    UserSaver
	usrProvider UserProvider
	appProvider AppProdider
	tokenTTL    time.Duration
}

type UserSaver interface {
	SaveUser(
		ctx context.Context,
		email string,
		passhash []byte,
	) (uid int64, err error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (models.User, error)
	IsAdmin(ctx context.Context, userId int64) (bool, error)
}
type AppProdider interface {
	App(ctx context.Context, appID int) (models.App, error)
}

func New(
	log *slog.Logger,
	userSaver UserSaver,
	userProvider UserProvider,
	appProvider AppProdider,
	tokenTTL time.Duration) *Auth {
	return &Auth{
		usrSaver:    userSaver,
		usrProvider: userProvider,
		log:         log,
		appProvider: appProvider,
		tokenTTL:    tokenTTL,
	}
}

func (a *Auth) Login(ctx context.Context, email string, password string, appID int,
) (string, error) {
	const op = "auth.Login"
	log := a.log.With(
		slog.String("op", op),
		slog.String("username", email))

	user, err := a.usrProvider.User(ctx, email)
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

	app, err := a.appProvider.App(ctx, appID)
	if err != nil {
		return "", fmt.Errorf("%s:%w", op, err)
	}
	log.Info("user logged successfully")

	token, err := jwt.NewToken(user, app, a.tokenTTL)
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
	id, err := a.usrSaver.SaveUser(ctx, email, passHash)
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
	isAdmin, err := a.usrProvider.IsAdmin(ctx, userId)
	if err != nil {
		if errors.Is(err, storage.ErrAppNotFound) {
			log.Error("failed to check if user is admin", sl.Err(err))
			return false, fmt.Errorf("%s : %w", op, err)
		}
	}
	log.Info("checked if the user is admin", slog.Bool("is_admin", isAdmin))
	return isAdmin, nil
}
