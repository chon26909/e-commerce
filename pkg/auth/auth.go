package auth

import "github.com/chon26909/e-commerce/config"

type TokenType string

const (
	Access  TokenType = "access"
	Refresh TokenType = "refresh"
	Admin   TokenType = "admin"
	ApiKey  TokenType = "apiKey"
)

type Auth struct {
	mapClaims *mapClaims
	config    config.IJwtConfig
}

type mapClaims struct {
}
