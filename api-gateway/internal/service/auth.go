package service

import (
	"context"
	"errors"

	"github.com/Yeprem25/go-auth-microservice/api-gateway/internal/model"
	"github.com/Yeprem25/go-auth-microservice/api-gateway/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	users     repository.UserRepository
	jwtSecret []byte
}

var ErrInvalidCredentials = errors.New("invalid credentials")

func NewAuthService(users repository.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{users: users, jwtSecret: []byte(jwtSecret)}
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

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.users.GetByEmail(ctx, email)

	if errors.Is(err, repository.ErrUserNotFound) {
		return "", ErrInvalidCredentials
	}
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", ErrInvalidCredentials
	}

	return s.generateAccessToken(user.ID)
}
