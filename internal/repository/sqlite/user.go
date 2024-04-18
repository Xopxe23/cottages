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

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	q := "SELECT * FROM users WHERE email=$1"
	var user domain.User
	err := r.db.QueryRow(q, email).Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Password, &user.IsActive, &user.IsVerified, &user.IsStaff)
	return user, err
}
