package utils

import (
	"blog/core/models"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTPayload struct {
	UserID           string `json:"user_id"`
	Email            string `json:"email"`
	RegisteredClaims jwt.StandardClaims
	jwt.Claims
}

func SignJWT(user *models.User) (string, error) {
	claims := JWTPayload{
		UserID: "12345",
		Email:  "test@example.com",
		RegisteredClaims: jwt.StandardClaims{
			ExpiresAt: int64(30 * 24 * time.Hour),
		},
	}

	jwtSecretKey := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

func VerifyJWT(tokenString string) (string, error) {
	jwtSecretKey := os.Getenv("JWT_SECRET")

	token, err := jwt.ParseWithClaims(tokenString, &JWTPayload{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			errorStr := fmt.Sprintf("unexpected signing method: %v", token.Header["alg"])
			return nil, errors.New(errorStr)
		}
		return []byte(jwtSecretKey), nil
	})

	if claims, ok := token.Claims.(*JWTPayload); ok && token.Valid {
		return claims.UserID, nil
	} else {
		return "", err
	}
}
