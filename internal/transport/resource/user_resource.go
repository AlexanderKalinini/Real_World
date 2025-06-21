package resource

import (
	"rwa/internal/domain/user"
	"time"
)

type WrapperResource struct {
	User UserResource `json:"user"`
}
type UserResource struct {
	Username  string
	Email     string
	Bio       string `json:"bio,omitempty"`
	Token     string `json:"token,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

func ConvertDomainToResource(user user.User) WrapperResource {
	return WrapperResource{
		User: UserResource{
			Username:  user.Username,
			Email:     user.Email,
			Bio:       user.Bio,
			Token:     user.Token,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		},
	}
}
