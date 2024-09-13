package response

import "api-auth/utils/models"

type UserLoginPayload struct {
	User  *models.User
	Token string
}
