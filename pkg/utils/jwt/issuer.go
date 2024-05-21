// Package jwt provides functionality for working with JSON Web Tokens (JWT).
//
// This package includes an Issuer type for creating JWT tokens using a private key.
package jwt

import (
	"OLO-backend/pkg/model"
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

func (i *Issuer) NewToken(user model.TokenUser, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString(i.key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
