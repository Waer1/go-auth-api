package helpers

import (
	appconstant "api-auth/utils/app-constant"
	"api-auth/utils/structs"
	"errors"

	"github.com/gin-gonic/gin"
)

// GetCurrentUser extracts the user from the Gin context and returns it as *UserJWT
func GetCurrentUser(ctx *gin.Context) (*structs.UserJWT, error) {
	// Extract the user object from the context
	user, exists := ctx.Get(appconstant.HeaderConstant.User)
	if !exists {
		return nil, errors.New("user not found in context")
	}

	// Type assertion to convert the interface{} to *structs.UserJWT
	userJWT, ok := user.(*structs.UserJWT)
	if !ok {
		return nil, errors.New("invalid user type in context")
	}

	return userJWT, nil
}
