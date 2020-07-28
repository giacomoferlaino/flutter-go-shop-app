package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func newJwtManager(signingKey string) *jwtManager {
	return &jwtManager{
		signingKey: []byte(signingKey),
		expiredIn:  24 * 60 * 60,
	}
}

type jwtManager struct {
	signingKey []byte
	expiredIn  uint64
}

func (manager *jwtManager) createToken() (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * time.Duration(manager.expiredIn)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(manager.signingKey)
}
