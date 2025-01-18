package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"
)

type JWTMiddleware struct {
	JWT *auth.JWTHandler
}

type JWTClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func NewJWTMiddleware(jwtSecret []byte) *JWTMiddleware {
	return &JWTMiddleware{
		JWT: auth.NewJWTHandler(jwtSecret),
	}
}

func (m *JWTMiddleware) HandleProtected(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header.
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		// Split the header into parts.
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			http.Error(w, "Invalid Authorization header", http.StatusBadRequest)
			return
		}

		// Get the JWT token from the header.
		jwtToken := parts[1]
		if jwtToken == "" {

			http.Error(w, "JWT token is missing", http.StatusUnauthorized)
			return
		}

		// Verify and validate the token.
		claims, err := m.JWT.ParseClaims(jwtToken)
		if err != nil {
			http.Error(w, "Invalid JWT token", http.StatusUnauthorized)
			return
		}

		// Check if the token is expired.
		if claims.Expiry.Before(time.Now()) {
			http.Error(w, "JWT token is expired", http.StatusUnauthorized)
			return
		}

		// Create a new claims struct.
		claimsStruct := &JWTClaims{
			UserID:   claims.UserID,
			Username: claims.Username,
			Email:    claims.Email,
		}

		// Create a new context.
		ctx := context.WithValue(r.Context(), "claims", claimsStruct)

		// Add the context to the request.
		r = r.WithContext(ctx)

		// Next.
		next.ServeHTTP(w, r)
	})
}

func (m *JWTMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			http.Error(w, "Invalid Authorization header", http.StatusBadRequest)
			return
		}

		jwtToken := parts[1]
		if jwtToken == "" {
			http.Error(w, "JWT token is missing", http.StatusUnauthorized)
			return
		}

		// Verify the token.
		claims, err := m.JWT.ParseClaims(jwtToken)
		if err != nil {
			http.Error(w, "Invalid JWT token", http.StatusUnauthorized)
			return
		}

		// Check if the token is expired.
		if claims.Expiry.Before(time.Now()) {
			http.Error(w, "JWT token is expired", http.StatusUnauthorized)
			return
		}

		// Send back the token.
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(jwtToken))
	}
}
