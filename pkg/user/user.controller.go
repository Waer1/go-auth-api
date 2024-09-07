package user

type UserController struct {
	userService UserService
}

func NewUserController(userService UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}
