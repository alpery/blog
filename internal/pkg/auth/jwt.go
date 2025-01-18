package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTHandler struct {
	Secret []byte
}

type JWTClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Expiry   time.Time `json:"expiry"`
	//TODO: add other data related to the user, e.g., profile picture
}

func NewJWTHandler(secret []byte) *JWTHandler {
	return &JWTHandler{Secret: secret}
}

func (j *JWTHandler) Generate(user *models.User) (string, error) {
	claims := JWTClaims{
		UserID:   user.ID.Hex(),
		Username: user.Username,
		Email:    user.Email,
		Expiry: time.Now().Add(time.Minute * 60), // 1 hour
	}

	token := jwt.NewWithClaims(claims, jwt.MapClaims{
		"user_id":   user.ID.Hex(),
		"username":  user.Username,
		"email":     user.Email,
		"expiry":    claims.Expiry,
	})

	// Sign the token using the secret key.
	token.Header.Set("typ", "JWT")
	token.Header.Set("alg", "HS256")
	token.Header.Set("kid", "default")

	// Sign the token.
	tokenString, err := token.SignedString(j.Secret)
	if err != nil {
		return "", fmt.Errorf("Failed to sign token: %v", err)
	}

	return tokenString, nil
}

func (j *JWTHandler) ParseClaims(tokenString string) (*JWTClaims, error) {
	// Parse the token.
	token, err := jwt.Parse(tokenString, j.Secret)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse token: %v", err)
	}

	// Get the claims from the token.
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, fmt.Errorf("Failed to get claims: %v", err)
	}

	return claims, nil
}
