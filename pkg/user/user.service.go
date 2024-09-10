package user

import (
	"api-auth/pkg/auth/dto"
	"api-auth/pkg/user/models"
	"api-auth/utils"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	Create(user *models.User) error
	FindUserBy(user *models.User) (*models.User, error)
	Login(loginDTO dto.LoginDTO) (*models.User, error)
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

// HashPassword hashes the given password using bcrypt
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (s *userService) Create(user *models.User) error {
	userWithSameEmail := models.User{Email: user.Email}
	existingUser, err := s.FindUserBy(&userWithSameEmail)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if existingUser != nil {
		return utils.NewServiceErr(http.StatusUnprocessableEntity, map[string]string{"email": "Email already exists"})
	}

	// Hash the user's password before saving
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return s.db.Create(user).Error
}

func (s *userService) FindUserBy(user *models.User) (*models.User, error) {
	var foundUser models.User
	if err := s.db.Where(user).First(&foundUser).Error; err != nil {
		return nil, err
	}
	return &foundUser, nil
}

func (s *userService) Login(loginDTO dto.LoginDTO) (*models.User, error) {
	userWithSameEmail := models.User{Email: loginDTO.Email}
	foundUser, err := s.FindUserBy(&userWithSameEmail)
	if err != nil {
		return nil, utils.NewServiceErr(http.StatusUnauthorized, map[string]string{"password": "Invalid Email or password"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(loginDTO.Password)); err != nil {
		return nil, utils.NewServiceErr(http.StatusUnauthorized, map[string]string{"password": "Invalid Email or password"})
	}
	return foundUser, nil
}
