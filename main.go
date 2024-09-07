package main

import (
	"api-auth/config"
	"api-auth/middleware"
	"api-auth/pkg/auth"
	"api-auth/pkg/user"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadConfig()
	config.ConnectDatabase()
}

func main() {
	// Initialize a new Gin router
	r := gin.Default()

	// Attach the error handling middleware
	r.Use(middleware.ErrorHandlingMiddleware())

	// Service initialization
	userService := user.NewUserService(config.DB)

	// Controller initialization
	authController := auth.NewAuthController(userService)

	// Authentication routes
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", authController.RegisterUser)
	}

	// Start the server on port 8080
	err := r.Run(":" + config.Config.AppPort)
	if err != nil {
		fmt.Println("Failed to start the server:", err)
	}
}
