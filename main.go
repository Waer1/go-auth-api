package main

import (
	"api-auth/config"
	"api-auth/middleware"
	"api-auth/pkg/auth"
	"api-auth/pkg/post"
	"api-auth/pkg/tag"
	"api-auth/pkg/user"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadConfig()
	config.ConnectDatabase()
	config.InitializeRedis()
}

func main() {
	// Initialize a new Gin router
	r := gin.Default()

	// Attach the error handling middleware
	r.Use(middleware.ErrorHandlingMiddleware())

	// Service initialization
	userService := user.NewUserService(config.DB)
	authService := auth.NewAuthService(userService, config.RedisClient)
	tagService := tag.NewTagService(config.DB)
	postService := post.NewPostService(config.DB, tagService)

	// Controller initialization
	authController := auth.NewAuthController(authService)
	tagController := tag.NewTagController(tagService)
	postController := post.NewPostController(postService)

	// Authentication routes
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", authController.RegisterUser)
		authRoutes.POST("/login", authController.LoginUser)
	}

	protectedRoutes := r
	protectedRoutes.Use(middleware.JWTAuthMiddleware(authService)) // Apply JWT middleware to the /api group
	{
		protectedRoutes.GET("/me", authController.Me) // Add the /me route

		// tags
		tagGroup := protectedRoutes.Group("/tags")
		tagGroup.POST("", tagController.CreateTag)
		tagGroup.GET("", tagController.GetAllTags)
		tagGroup.PATCH("/:id", tagController.UpdateTag)
		tagGroup.DELETE("/:id", tagController.DeleteTag)
		tagGroup.GET("/:id", tagController.GetTag)

		// posts
		postGroup := protectedRoutes.Group("/posts")
		postGroup.POST("/", postController.CreatePost)
		postGroup.GET("/", postController.GetAllPosts)
		postGroup.PATCH("/:id", postController.UpdatePost)
		postGroup.DELETE("/:id", postController.DeletePost)
		postGroup.GET("/:id", postController.GetPostById)

	}

	// Start the server on port 8080
	err := r.Run(":" + config.Config.AppPort)
	if err != nil {
		fmt.Println("Failed to start the server:", err)
	}
}
