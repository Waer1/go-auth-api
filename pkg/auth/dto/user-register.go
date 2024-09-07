package dto

// UserRegistrationDTO is used when registering a new user
type UserRegistrationDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
