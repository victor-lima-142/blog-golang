package middlewares

import (
	"blog/src/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func LoggedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := godotenv.Load()

		if err != nil {
			log.Fatalf("Error loading .env file")
		}

		jwtSecretKey := os.Getenv("JWT_SECRET")

		authorization := r.Header.Get("Authorization")

		hasPrefix := strings.HasPrefix(authorization, "Bearer ")
		if (authorization == "") || (!hasPrefix) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		authHeader := authorization
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			http.Error(w, "Authorization header must be in Bearer token format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		token, err := jwt.ParseWithClaims(tokenString, &utils.JWTPayload{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecretKey, nil
		})

		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid token: %v", err), http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(*utils.JWTPayload)
		if !ok || !token.Valid {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		if claims.UserID == "" || claims.Email == "" {
			http.Error(w, "Token claims are incomplete", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
