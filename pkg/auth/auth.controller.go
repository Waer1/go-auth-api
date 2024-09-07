package auth

import (
	"api-auth/pkg/auth/dto"
	"api-auth/pkg/user"
	"api-auth/pkg/user/models"
	"api-auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService user.UserService
}

func NewAuthController(userService user.UserService) AuthController {
	return AuthController{userService: userService}
}

func (ac *AuthController) RegisterUser(c *gin.Context) {
	var userDTO dto.UserRegistrationDTO
	if !utils.BindJSONAndValidate(c, &userDTO) {
		return
	}

	user := models.User{
		Email:    userDTO.Email,
		Password: userDTO.Password, // Password hashing should be handled inside the service layer
	}

	err := ac.userService.Create(&user)
	if err != nil {
		c.Error(err) // Add error to Gin Context
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// how to prevent the server to run if there is any error and show that error
