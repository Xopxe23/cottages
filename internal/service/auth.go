package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/Xopxe23/cottages/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user domain.User) error
	GetByEmail(ctx context.Context, email string) (domain.User, error)
}

type PasswordHasher interface {
	Hash(password string) (string, error)
	VerifyPassword(password, hashedPassword string) error
}

type TokenRepository interface {
	Create(ctx context.Context, token domain.RefreshToken) error
	Get(ctx context.Context, token string) (domain.RefreshToken, error)
}

type AuthService struct {
	userRepository  UserRepository
	tokenRepository TokenRepository
	hasher          PasswordHasher

	hmacSecret []byte
	tokenTTL   time.Duration
}

func NewAuthService(userRepo UserRepository, tokenRepo TokenRepository, hasher PasswordHasher, secret []byte, ttl time.Duration) *AuthService {
	return &AuthService{userRepo, tokenRepo, hasher, secret, ttl}
}

func (s *AuthService) CreateUser(ctx context.Context, input domain.SignUpInput) error {
	email := strings.ToLower(*input.Email)
	password, err := s.hasher.Hash(*input.Password)
	if err != nil {
		return err
	}
	user := domain.User{
		ID:        uuid.New().String(),
		Email:     email,
		FirstName: *input.FirstName,
		LastName:  *input.LastName,
		Password:  password,
	}
	return s.userRepository.Create(ctx, user)
}

func (s *AuthService) Authenticate(ctx context.Context, input domain.SignInInput) (string, string, error) {
	email := strings.ToLower(*input.Email)
	user, err := s.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return "", "", err
	}
	if err := s.hasher.VerifyPassword(*input.Password, user.Password); err != nil {
		return "", "", err
	}
	accessToken, refreshToken, err := s.generateTokens(ctx, user.ID)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, err
}

func (s *AuthService) generateTokens(ctx context.Context, userId string) (string, string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":       userId,
		"issuedAt":  time.Now().Unix(),
		"expiredAt": time.Now().Add(s.tokenTTL).Unix(),
	})
	accessToken, err := t.SignedString(s.hmacSecret)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := newRefreshToken()
	if err != nil {
		return "", "", err
	}
	token := domain.RefreshToken{
		ID:        uuid.New().String(),
		UserID:    userId,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30),
	}
	if err = s.tokenRepository.Create(ctx, token); err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func newRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

func (s *AuthService) RefreshTokens(ctx context.Context, refreshToken string) (string, string, error) {
	refresh, err := s.tokenRepository.Get(ctx, refreshToken)
	if err != nil {
		return "", "", err
	}
	if refresh.ExpiresAt.Unix() < time.Now().Unix() {
		return "", "", errors.New("refreshToken expired")
	}
	return s.generateTokens(ctx, refresh.UserID)
}
