package request

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"rwa/internal/domain/user"
	"rwa/pkg/crypto/argon"
	"rwa/pkg/session"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		db: db,
	}
}

func (r Repository) Authenticate(ctx context.Context, token string) (*user.User, error) {

	row := r.db.QueryRowContext(ctx, "SELECT email, username,created_at,updated_at FROM users WHERE token = ?", token)
	var dbUser user.User
	err := row.Scan(&dbUser.Email, &dbUser.Username, &dbUser.CreatedAt, &dbUser.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("не валидный токен")
		}
		return nil, err
	}

	return &dbUser, nil
}

func (r Repository) Login(ctx context.Context, u *user.User) (*user.User, error) {

	row := r.db.QueryRow("SELECT username,email,created_at,updated_at, COALESCE(token,''), COALESCE(bio,''),password FROM users WHERE email = ?", u.Email)
	var dbUser user.User
	err := row.Scan(
		&dbUser.Username,
		&dbUser.Email,
		&dbUser.CreatedAt,
		&dbUser.UpdatedAt,
		&dbUser.Token,
		&dbUser.Bio,
		&dbUser.Password,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("учётные данные не верны")
		}
		return nil, fmt.Errorf("scan error: %w", err)
	}

	token, err := session.GenerateSessionToken()

	if err != nil {
		return nil, fmt.Errorf("генерация сессии вернула ошибку %w", err)
	}

	dbUser.Token = token

	_, err = r.db.ExecContext(ctx, "UPDATE users SET token = ? WHERE email=?", dbUser.Token, dbUser.Email)

	if err != nil {
		return nil, err
	}

	if !argon.CheckPass(u.Password, []byte(dbUser.Password)) {
		return nil, fmt.Errorf("учётные данные не верны")
	}
	return &dbUser, nil
}

func (r Repository) Create(ctx context.Context, user *user.User) error {

	var count int
	err := r.db.QueryRow("SELECT COUNT(email) FROM users WHERE email = ?", user.Email).Scan(&count)

	if err != nil {
		return err
	}

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

	_, err = stmt.ExecContext(ctx, user.Username, user.Email, argon.GetHashPass(user.Password, nil))
	if err != nil {
		return fmt.Errorf("ошибка выполнения запроса: %w", err)
	}

	return nil
}
