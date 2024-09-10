package auth

import (
	"api-auth/pkg/auth/dto"
	"api-auth/pkg/user/models"
	"api-auth/utils"
	"api-auth/utils/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService AuthService
}

func NewAuthController(authService AuthService) AuthController {
	return AuthController{authService: authService}
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

	err := ac.authService.RegisterUser(&user)
	if err != nil {
		c.Error(err) // Add error to Gin Context
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (ac *AuthController) LoginUser(c *gin.Context) {
	var loginDTO dto.LoginDTO
	if !utils.BindJSONAndValidate(c, &loginDTO) {
		return
	}

	data, err := ac.authService.LoginUser(loginDTO)
	if err != nil {
		c.Error(err) // Add error to Gin Context
		return
	}

	c.JSON(http.StatusOK, data)
}

func (ac *AuthController) Me(c *gin.Context) {
	user, err := helpers.GetCurrentUser(c)
	if err != nil {
		c.Error(err) // Add error to Gin Context
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}
