package auth

import (
	"math"
	"time"

	"github.com/chon26909/e-commerce/config"
	"github.com/chon26909/e-commerce/modules/users"
	"github.com/golang-jwt/jwt/v5"
)

type TokenType string

const (
	Access  TokenType = "access"
	Refresh TokenType = "refresh"
	Admin   TokenType = "admin"
	ApiKey  TokenType = "apiKey"
)

type auth struct {
	mapClaims *MapClaims
	config    config.IJwtConfig
}

type MapClaims struct {
	Claims *users.UserClaims `json:"claims"`
	jwt.RegisteredClaims
}

type IAuth interface {
	SignToken() string
}

func jwtTimeDuration(t int) *jwt.NumericDate {

	return jwt.NewNumericDate(time.Now().Add(time.Duration(int64(t) * int64(math.Pow10(9)))))
}

func jwtTimeRepeatAdaptre(t int64) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Unix(t, 0))
}

func NewAuth(tokenType string, config config.IConfig, claims *users.UserClaims) (IAuth, error) {

	switch tokenType {
	case string(Access):
		return newAccessToken(config, claims), nil
	}

	return nil, nil
}

func newAccessToken(config config.IJwtConfig, claims *users.UserClaims) IAuth {
	return &auth{
		config: config,
		mapClaims: &MapClaims{
			Claims: claims,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "shop-api",
				Subject:   "access-token",
				Audience:  []string{"customer", "admin"},
				ExpiresAt: jwtTimeDuration(config.AccessExpiresAt()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		},
	}
}
