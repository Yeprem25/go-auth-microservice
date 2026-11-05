package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"

	"github.com/Yeprem25/go-auth-microservice/api-gateway/internal/model"
	"github.com/Yeprem25/go-auth-microservice/api-gateway/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	users     repository.UserRepository
	tokens    repository.RefreshTokenRepository
	jwtSecret []byte
}

var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrInvalidRefreshToken = errors.New("invalid refresh token")

func NewAuthService(users repository.UserRepository, tokens repository.RefreshTokenRepository, jwtSecret string) *AuthService {
	return &AuthService{users: users, tokens: tokens, jwtSecret: []byte(jwtSecret)}
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

func (s *AuthService) Login(ctx context.Context, email, password string) (string, string, error) {
	user, err := s.users.GetByEmail(ctx, email)

	if errors.Is(err, repository.ErrUserNotFound) {
		return "", "", ErrInvalidCredentials
	}
	if err != nil {
		return "", "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", "", ErrInvalidCredentials
	}

	accessToken, err := s.generateAccessToken(user.ID)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.createRefreshToken(ctx, user.ID)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) GetUser(ctx context.Context, id int64) (*model.User, error) {
	return s.users.GetByID(ctx, id)
}

func (s *AuthService) createRefreshToken(ctx context.Context, userID int64) (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	token := hex.EncodeToString(b)

	err := s.tokens.Create(ctx, &model.RefreshToken{
		UserID:    userID,
		TokenHash: hashToken(token),
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	})
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) Logout(ctx context.Context, refreshToken string) error {
	rt, err := s.tokens.GetByHash(ctx, hashToken(refreshToken))
	if errors.Is(err, repository.ErrRefreshTokenNotFound) {
		return nil
	}
	if err != nil {
		return err
	}

	return s.tokens.Delete(ctx, rt.ID)
}
