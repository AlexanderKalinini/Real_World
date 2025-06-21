package user

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, user *User) error
	Login(ctx context.Context, user *User) (*User, error)
	Authenticate(ctx context.Context, token string) (*User, error)
}
