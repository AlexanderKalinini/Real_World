package user

import (
	"database/sql"
)

//type UserRepository interface {
//	Get(ctx context.Context, db sql.DB, user user.User)
//	Create(ctx context.Context, db sql.DB, user user.User)
//}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}
