package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT generates a JWT token with a given payload, secret, and expiration duration
func GenerateJWT(payload any, secret []byte, expiration time.Duration) (string, error) {
	// Set expiration time
	expirationTime := time.Now().Add(expiration)

	// Create claims
	claims := jwt.MapClaims{
		"exp": expirationTime.Unix(),
	}

	// Add the payload to claims
	claims["payload"] = payload

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseJWT parses a JWT token and returns the payload
func ValidateToken(tokenString string, claims jwt.Claims, secret string) (interface{}, error) {
	// Parse the token with the provided claims and secret
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	// Return the claims as interface{}
	return claims, nil
}
