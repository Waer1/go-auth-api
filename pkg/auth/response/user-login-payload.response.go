package response

import "api-auth/pkg/user/models"

type UserLoginPayload struct {
	User  *models.User
	Token string
}
