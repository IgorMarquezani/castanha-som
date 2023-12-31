package userfeatures

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) (*repository, error) {
	if db == nil {
		return nil, errors.New("db pointer cannot be nil")
	}

	return &repository{db: db}, nil
}

func (r *repository) Create(ctx context.Context, features *UserFeatures) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	return r.db.WithContext(ctx).Create(features).Error
}

func (r *repository) SelectByUserID(ctx context.Context, userID string) (UserFeatures, error) {
	var features UserFeatures

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&features).Error

	return features, err
}
