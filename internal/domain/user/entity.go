package user

import "time"

type User struct {
	Username  string
	Email     string
	Bio       string
	Token     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
