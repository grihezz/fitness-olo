package jwt

import (
	"OLO-backend/auth_service/internal/domain/models"
	"crypto"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type Issuer struct {
	key crypto.PrivateKey
}

func NewIssuer(privateKeyPath string) (*Issuer, error) {
	keyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		panic(fmt.Errorf("unable to read private key file: %w", err))
	}

	key, err := jwt.ParseEdPrivateKeyFromPEM(keyBytes)
	if err != nil {
		return nil, fmt.Errorf("unable to parse as ed private key: %w", err)
	}

	return &Issuer{
		key: key,
	}, nil
}

func (i *Issuer) NewToken(user *models.User, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString(i.key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
