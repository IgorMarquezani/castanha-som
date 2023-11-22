package cart

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) (*Repository, error) {
	if db == nil {
		return nil, errors.New("db pointer cannot be nil")
	}

	return &Repository{db: db}, nil
}

func (r *Repository) Create(ctx context.Context, cart *Cart) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	return r.db.WithContext(ctx).Create(cart).Error
}

func (r *Repository) SelectByUser(ctx context.Context, userID string) (Cart, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	var cart Cart

	err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&cart).Error

	return cart, err
}
