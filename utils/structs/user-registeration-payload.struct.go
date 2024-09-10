package structs

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

// UserJWT represents the custom claims structure for the user.
type UserJWT struct {
	UserId uint
	Email  string
}

// DecodeUserJWT decodes the user information from jwt.MapClaims.
func DecodeUserJWT(claims jwt.MapClaims) (*UserJWT, error) {
	// Extract the payload from the claims
	payload, ok := claims["payload"].(map[string]interface{})
	if !ok {
		return nil, errors.New("invalid payload format")
	}

	fmt.Printf("Payload: %v\n", payload)

	// Extract and convert the UserId from float64 to uint
	userIdFloat, ok := payload["UserId"].(float64) // JSON numbers are parsed as float64
	if !ok {
		return nil, errors.New("invalid UserId format")
	}
	userId := uint(userIdFloat) // Convert float64 to uint

	// Extract Email
	email, ok := payload["Email"].(string)
	if !ok {
		return nil, errors.New("invalid Email format")
	}

	// Return the populated UserJWT struct
	return &UserJWT{
		UserId: userId,
		Email:  email,
	}, nil
}
