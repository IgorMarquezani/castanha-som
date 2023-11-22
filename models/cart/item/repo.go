package item

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

func (r *Repository) Create(ctx context.Context, item *Item) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	return r.db.WithContext(ctx).Table("cart_items").Create(item).Error
}

func (r *Repository) SelectByCarID(ctx context.Context, cartID string) ([]Item, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

  items := make([]Item, 0, 0)

  err := r.db.WithContext(ctx).Table("cart_items").Where("cart_id = ?", cartID).Find(&items).Error
  
  return items, err
}
