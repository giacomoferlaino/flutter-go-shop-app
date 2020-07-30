package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type customClaims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

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
func (manager *JwtManager) CreateToken(id uint) (string, error) {
	expiresAt := time.Now().Add(time.Second * time.Duration(manager.ExpiredIn)).Unix()
	claims := &customClaims{id, jwt.StandardClaims{ExpiresAt: expiresAt}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(manager.SigningKey)
}

// ParseToken determines if a token is valid and if it's not it returns an error
func (manager *JwtManager) ParseToken(token string) (uint, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return manager.SigningKey, nil
	})
	if err != nil {
		return 0, errors.New("Invalid authentication token")
	}

	claims, ok := parsedToken.Claims.(*customClaims)
	if !ok || !parsedToken.Valid {
		return 0, errors.New("Invalid authentication token")
	}
	return claims.ID, nil
}
