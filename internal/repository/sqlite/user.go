package repository

import (
	"context"
	"database/sql"

	"github.com/Xopxe23/cottages/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(ctx context.Context, user domain.User) error {
	q := "INSERT INTO users(id, email, first_name, last_name, password) VALUES ($1, $2, $3, $4, $5)"
	_, err := r.db.Exec(q, user.ID, user.Email, user.FirstName, user.LastName, user.Password)
	return err
}
