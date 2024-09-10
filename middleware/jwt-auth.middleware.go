package middleware

import (
	"api-auth/pkg/auth"
	appconstant "api-auth/utils/app-constant"
	"api-auth/utils/structs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWTAuthMiddleware validates the JWT token and injects the user into the context.
func JWTAuthMiddleware(authService auth.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
			ctx.Abort()
			return
		}

		tokenString := authHeader[len("Bearer "):]

		// Define the claims struct
		claims := jwt.MapClaims{}

		// Validate the token using the AuthService
		_, err := authService.ValidateToken(tokenString, claims)
		if err != nil {
			fmt.Println("Error validating token:", err)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			ctx.Abort()
			return
		}

		// Decode the UserJWT from the claims
		userJWT, err := structs.DecodeUserJWT(claims)
		if err != nil {
			fmt.Println("Error decoding user JWT:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode JWT"})
			ctx.Abort()
			return
		}

		// Inject the entire userJWT object into the context
		ctx.Set(appconstant.HeaderConstant.User, userJWT)

		ctx.Next() // Continue to the next handler
	}
}
