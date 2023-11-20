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

func (r *Repository) Select(ctx context.Context) ([]Product, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	products := make([]Product, 0, 10)

	err := r.db.WithContext(ctx).Find(&products).Error

	return products, err
}

func (r *Repository) SelectByType(ctx context.Context, Type string) ([]Product, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	products := make([]Product, 0, 10)

	err := r.db.WithContext(ctx).Where("type = ?", Type).Find(&products).Error

	return products, err
}

func (r *Repository) Delete(ctx context.Context, productName string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	return r.db.WithContext(ctx).Where("name = ?", productName).Delete(&Product{}).Error
}

func (r *Repository) RawSelect(ctx context.Context, rawSQL string, values ...any) ([]Product, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	products := make([]Product, 0, 10)

	err := r.db.WithContext(ctx).Raw(rawSQL, values...).Find(&products).Error

	return products, err
}
