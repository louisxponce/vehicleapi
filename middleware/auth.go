package middleware

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const claimsKey contextKey = "jwtClaims"

func AuthMiddleware(publicKey *rsa.PublicKey) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			tokenStr := getBearerToken(r)
			if tokenStr == "" {
				unauthorized(w, "Missing or invalid Authorization header")
				return
			}

			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
				// Validate the signing method
				if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return publicKey, nil
			})

			if err != nil || !token.Valid {
				unauthorized(w, "Invalid token")
				return
			}

			// Claims
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				unauthorized(w, "Invalid token claims")
				return
			}

			ctx := context.WithValue(r.Context(), claimsKey, claims)
			next(w, r.WithContext(ctx))
		}
	}
}

func getBearerToken(r *http.Request) string {
	auth := r.Header.Get("Authorization")
	if strings.HasPrefix(auth, "Bearer ") {
		return strings.TrimPrefix(auth, "Bearer ")
	}
	return ""
}

func unauthorized(w http.ResponseWriter, msg string) {
	log.Printf("Unauthorized")
	w.Header().Set("content_type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(map[string]string{
		"error":             "invalid_client",
		"error_description": msg,
	})
}

func GetClaims(r *http.Request) jwt.MapClaims {
	val := r.Context().Value(claimsKey)
	if claims, ok := val.(jwt.MapClaims); ok {
		return claims
	}
	return nil
}
