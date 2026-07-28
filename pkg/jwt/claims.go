package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/motorcyclists-network/go-common/pkg/enum"
)

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

type CustomClaims struct {
	AccountID string    `json:"account_id"`
	Role      enum.Role `json:"role"`
	TokenType TokenType `json:"token_type"`
	jwt.RegisteredClaims
}
