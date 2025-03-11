package user

import (
	"context"
	"database/sql"
	"fmt"
	"rwa/internal/model/user"
	"rwa/pkg/crypto/argon"
)

func (r *Repository) Create(ctx context.Context, user *user.User) error {

	var count int
	r.db.QueryRow("SELECT COUNT(email) FROM users WHERE email = ?", user.Email).Scan(&count)

	if count > 0 {
		return fmt.Errorf("пользователь с таким email уже существует")
	}

	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("ошибка подготовки запроса: %w", err)

	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			return
		}
	}(stmt)

	_, err = stmt.ExecContext(ctx, user.Username, user.Email, argon.GetHashPass(nil, user.Password))
	if err != nil {
		return fmt.Errorf("ошибка выполнения запроса: %w", err)
	}

	return nil
}
