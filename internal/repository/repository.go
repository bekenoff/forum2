package repository

import (
	"database/sql"
	"forum/internal/repository/auth"
)

type Repository struct {
	auth.Auth
}

func NewRepo(db *sql.DB) *Repository {
	return &Repository{
		Auth: auth.NewAuthRepo(db),
	}
}
