package description

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Respository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) (*Respository, error) {
	if db == nil {
		return nil, errors.New("db pointer cannot be nil")
	}

	return &Respository{db: db}, nil
}

func (r *Respository) Create(ctx context.Context, description *Description) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	return r.db.WithContext(ctx).Table("product_descriptions").Create(description).Error
}

func (r *Respository) SelectByProduct(ctx context.Context, productName string) ([]Description, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	des := make([]Description, 0, 0)

	err := r.db.WithContext(ctx).Table("product_descriptions").Where("product_name = ?", productName).Find(&des).Error

	return des, err
}

func (r *Respository) DeleteByProductName(ctx context.Context, productName string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	return r.db.WithContext(ctx).Table("product_descriptions").Where("product_name = ?", productName).
		Delete(&Description{}).Error
}
