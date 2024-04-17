package service

import (
	"context"

	"github.com/Xopxe23/cottages/internal/domain"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user domain.User) error
}

type PasswordHasher interface {
	Hash(password string) (string, error)
}

type AuthService struct {
	userRepository UserRepository
	hasher         PasswordHasher
}

func NewAuthService(userRepo UserRepository, hasher PasswordHasher) *AuthService {
	return &AuthService{userRepo, hasher}
}

func (s *AuthService) CreateUser(ctx context.Context, input domain.SignUpInput) error {
	password, err := s.hasher.Hash(*input.Password)
	if err != nil {
		return err
	}
	user := domain.User{
		ID: uuid.New().String(),
		Email: *input.Email,
		FirstName: *input.FirstName,
		LastName: *input.LastName,
		Password: password,
	}
	return s.userRepository.Create(ctx, user)
}
