package auth

import (
	"errors"
	"fmt"
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
	config    config.IJwtConfig
	mapClaims *MapClaims
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

// func jwtTimeRepeatAdaptre(t int64) *jwt.NumericDate {
// 	return jwt.NewNumericDate(time.Unix(t, 0))
// }

func (a *auth) SignToken() string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a.mapClaims)

	signedToken, _ := token.SignedString(a.config.SecertKey())

	return signedToken
}

func ParseToken(config config.IJwtConfig, tokenString string) (*auth, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MapClaims{}, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid: %v", t.Method)
		}
		return config.SecertKey(), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {

		}
	}

	return nil, nil
}

func NewAuth(tokenType string, config config.IJwtConfig, claims *users.UserClaims) (IAuth, error) {

	switch tokenType {
	case string(Access):
		return newAccessToken(config, claims), nil
	case string(Refresh):
		return newRefreshToken(config, claims), nil
	default:
		return nil, fmt.Errorf("unknow token type")
	}
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

func newRefreshToken(config config.IJwtConfig, claims *users.UserClaims) IAuth {
	return &auth{
		config: config,
		mapClaims: &MapClaims{
			Claims: claims,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "shop-api",
				Subject:   "refresh-token",
				Audience:  []string{"customer", "admin"},
				ExpiresAt: jwtTimeDuration(config.RefreshExpiresAt()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		},
	}
}
