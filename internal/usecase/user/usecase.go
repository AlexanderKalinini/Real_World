package user

import (
	"context"
	"errors"
	"rwa/internal/domain/user"
)

type UseCase interface {
	Create(ctx context.Context, user *user.User) error
	Login(ctx context.Context, user *user.User) (*user.User, error)
	Authenticate(ctx context.Context, token string) (*user.User, error)
}

type useCase struct {
	repo user.Repository
}

func NewUseCase(r user.Repository) UseCase {
	return &useCase{repo: r}
}

func (u *useCase) Authenticate(ctx context.Context, token string) (*user.User, error) {
	return u.repo.Authenticate(ctx, token)
}

func (u *useCase) Create(ctx context.Context, user *user.User) error {
	if user.Email == "" || user.Password == "" {
		return errors.New("email and password are required")
	}

	return u.repo.Create(ctx, user)
}

func (u *useCase) Login(ctx context.Context, user *user.User) (*user.User, error) {
	if user.Email == "" || user.Password == "" {
		return nil, errors.New("email and password are required")
	}

	return u.repo.Login(ctx, user)
}
