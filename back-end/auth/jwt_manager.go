package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// NewJwtManager return a new jwt manager
func NewJwtManager(signingKey string) *JwtManager {
	return &JwtManager{
		SigningKey: []byte(signingKey),
		ExpiredIn:  24 * 60 * 60,
	}
}

// JwtManager is used to managed JWT for authentication
type JwtManager struct {
	SigningKey []byte
	ExpiredIn  uint64
}

// CreateToken returns a new JWT
func (manager *JwtManager) CreateToken() (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * time.Duration(manager.ExpiredIn)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(manager.SigningKey)
}
