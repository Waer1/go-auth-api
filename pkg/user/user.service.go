package user

import (
	"api-auth/pkg/user/models"
	"api-auth/utils"
	"net/http"

	"gorm.io/gorm"
)

type UserService interface {
	Create(user *models.User) error
	FindUserBy(user *models.User) (*models.User, error)
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
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

	return s.db.Create(user).Error
}

func (s *userService) FindUserBy(user *models.User) (*models.User, error) {
	var foundUser models.User
	if err := s.db.Where(user).First(&foundUser).Error; err != nil {
		return nil, err
	}
	return &foundUser, nil
}
