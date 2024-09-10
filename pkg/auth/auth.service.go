package auth

import (
	"api-auth/config"
	"api-auth/pkg/auth/dto"
	"api-auth/pkg/auth/response"
	"api-auth/pkg/user"
	"api-auth/pkg/user/models"
	"api-auth/utils"
	"api-auth/utils/structs"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	RegisterUser(user *models.User) error
	LoginUser(loginDTO dto.LoginDTO) (response.UserLoginPayload, error)
	ValidateToken(tokenString string, claims jwt.Claims) (any, error)
}

type authService struct {
	userService user.UserService
}

func NewAuthService(userService user.UserService) AuthService {
	return &authService{
		userService: userService,
	}
}

// LoginUser implements AuthService.
func (as *authService) LoginUser(loginDTO dto.LoginDTO) (response.UserLoginPayload, error) {
	user, err := as.userService.Login(loginDTO)
	if err != nil {
		return response.UserLoginPayload{}, err
	}

	jwtPayload := &structs.UserJWT{
		UserId: user.ID,
		Email:  user.Email,
	}

	userToken, err := utils.GenerateJWT(jwtPayload, []byte(config.Config.JwtSecret), config.Config.JwtExpireIn)
	if err != nil {
		return response.UserLoginPayload{}, err
	}

	return response.UserLoginPayload{
		User:  user,
		Token: userToken,
	}, nil
}

// RegisterUser implements AuthService.
func (as *authService) RegisterUser(user *models.User) error {
	return as.userService.Create(user)
}

// ValidateToken implements AuthService.
func (as *authService) ValidateToken(tokenString string, claims jwt.Claims) (any, error) {
	return utils.ValidateToken(tokenString, claims, config.Config.JwtSecret)
}
