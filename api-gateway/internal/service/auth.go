package service

import (
	"context"
	"errors"

	"github.com/Yeprem25/go-auth-microservice/api-gateway/internal/model"
	"github.com/Yeprem25/go-auth-microservice/api-gateway/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	users repository.UserRepository
}

var ErrInvalidCredentials = errors.New("invalid credentials")

func NewAuthService(users repository.UserRepository) *AuthService {
	return &AuthService{users: users}
}

func (s *AuthService) Register(ctx context.Context, email, password string) (*model.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &model.User{Email: email, PasswordHash: string(hash)}
	err = s.users.Create(ctx, user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(ctx context.Context, email, password string) (*model.User, error) {
	user, err := s.users.GetByEmail(ctx, email)

	if errors.Is(err, repository.ErrUserNotFound) {
		return nil, ErrInvalidCredentials
	}
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}
