package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/Yeprem25/go-auth-microservice/auth-service/internal/repository"
	"github.com/golang-jwt/jwt/v5"
)

func (s *AuthService) generateAccessToken(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(15 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecret)
}

func (s *AuthService) ParseAccessToken(tokenStr string) (int64, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		return s.jwtSecret, nil
	}, jwt.WithValidMethods([]string{"HS256"}))
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token")
	}

	sub, ok := claims["sub"].(float64)
	if !ok {
		return 0, errors.New("invalid token subject")
	}

	return int64(sub), nil
}

func (s *AuthService) Refresh(ctx context.Context, refreshToken string) (string, string, error) {
	rt, err := s.tokens.GetByHash(ctx, hashToken(refreshToken))
	if errors.Is(err, repository.ErrRefreshTokenNotFound) {
		return "", "", ErrInvalidRefreshToken
	}
	if err != nil {
		return "", "", err
	}

	if time.Now().After(rt.ExpiresAt) {
		return "", "", ErrInvalidRefreshToken
	}
	if err := s.tokens.Delete(ctx, rt.ID); err != nil {
		return "", "", err
	}

	accessToken, err := s.generateAccessToken(rt.UserID)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = s.createRefreshToken(ctx, rt.UserID)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func hashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}
