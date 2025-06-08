package user

import "rwa/internal/repository/user"

type UsersController struct {
	UserRepo *user.Repository
}

func NewUserController(userRepo *user.Repository) *UsersController {
	return &UsersController{
		UserRepo: userRepo,
	}
}
