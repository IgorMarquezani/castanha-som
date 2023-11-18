package product

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

func (r *Repository) Create(ctx context.Context, product *Product) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	return r.db.WithContext(ctx).Create(&product).Error
}

func (r *Repository) Delete(ctx context.Context, productName string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	return r.db.WithContext(ctx).Where("name = ?", productName).Delete(&Product{}).Error
}