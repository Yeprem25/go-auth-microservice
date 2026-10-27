package repository

import (
	"context"
	"errors"

	"github.com/Yeprem25/go-auth-microservice/api-gateway/internal/model"
	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	db *gorm.DB
}

var _ UserRepository = (*PostgresUserRepository)(nil)

func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(ctx context.Context, user *model.User) error {
	err := r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return ErrEmailTaken
		}
		return err
	}

	return nil
}

func (r *PostgresUserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}
