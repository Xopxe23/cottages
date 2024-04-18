package repository

import (
	"context"
	"database/sql"

	"github.com/Xopxe23/cottages/internal/domain"
)

type TokenRepository struct {
	db *sql.DB
}

func NewTokenRepository(db *sql.DB) *TokenRepository {
	return &TokenRepository{db}
}

func (r *TokenRepository) Create(ctx context.Context, token domain.RefreshToken) error {
	q := "INSERT INTO refresh_tokens (id, user_id, token, expires_at) VALUES($1, $2, $3, $4)"
	_, err := r.db.Exec(q, token.ID, token.UserID, token.Token, token.ExpiresAt)
	return err
}

func (r *TokenRepository) Get(ctx context.Context, token string) (domain.RefreshToken, error) {
	var t domain.RefreshToken
	q := "SELECT id, user_id, token, expires_at FROM refresh_tokens WHERE token=$1"
	if err := r.db.QueryRow(q, token).
		Scan(&t.ID, &t.UserID, &t.Token, &t.ExpiresAt); err != nil {
		return t, err
	}
	q = "DELETE FROM refresh_tokens WHERE user_id=$1"
	_, err := r.db.Exec(q, t.UserID)
	return t, err
}
