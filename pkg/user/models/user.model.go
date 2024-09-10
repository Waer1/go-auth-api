package models

import (
	"api-auth/utils/models"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	models.BaseEntity
	Email    string `gorm:"uniqueIndex;size:100" json:"email"` // Ensure email is unique and define a reasonable size
	Password string `json:"-"`
}

// HashPassword hashes the user's password using bcrypt.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // Use DefaultCost instead of hard coding
	return string(bytes), err
}
