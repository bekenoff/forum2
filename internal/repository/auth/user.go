package auth

import (
	"database/sql"
	"forum/internal/models"
)

type AuthRepo struct {
	DB *sql.DB
}

type Auth interface {
	CreateClient(client models.UserRegistration) error
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{
		DB: db,
	}
}

func (u *AuthRepo) SaveUser(user models.UserRegistration) error {
	return nil
}

func (u *AuthRepo) CreateClient(client models.UserRegistration) error {
	stmt := `
	INSERT INTO users (name, email, password)
	VALUES (?, ?, ?);
	`
	_, err := u.DB.Exec(stmt, client.Username, client.Email, client.Password)
	if err != nil {
		return err
	}
	return nil
}
