package user

import "rwa/internal/repository/user"

type CreateUserController struct {
	UserRepo *user.Repository
}

func NewUserController(userRepo *user.Repository) *CreateUserController {
	return &CreateUserController{
		UserRepo: userRepo,
	}
}
