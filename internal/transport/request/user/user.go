package user

import (
	"rwa/internal/domain/user"
	"time"
)

type User struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email" validate:"required,email"`
	Bio      string `json:"bio" validate:"max=20000"`
	Token    string `json:"token"`
	Password string `json:"password" validate:"required"`
}

type UsersWrapper struct {
	Users User `json:"user"`
}

func ConvertUserToDomain(requestUser User) user.User {
	return user.User{
		Username:  requestUser.Username,
		Email:     requestUser.Email,
		Bio:       requestUser.Bio,
		Token:     requestUser.Token,
		Password:  requestUser.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
