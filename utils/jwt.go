package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(sub uint, secret string, expHours int) (string, error) {
	claims := jwt.MapClaims{
		"sub": sub,
		"exp": time.Now().Add(time.Duration(expHours) * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(secret))
}
