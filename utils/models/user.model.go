package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseEntity
	Email    string `gorm:"uniqueIndex;size:100" json:"email"` // Ensure email is unique and define a reasonable size
	Password string `json:"-"`
	Posts    []Post `json:"posts"`
}

// HashPassword hashes the user's password using bcrypt.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // Use DefaultCost instead of hard coding
	return string(bytes), err
}
