package user

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

	return &repository{
		db: db,
	}, nil
}

func (u *repository) SelectByID(ctx context.Context, userId string) (User, error) {
	var user User

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	return user, u.db.WithContext(ctx).Where("id = ?", userId).First(&user).Error
}

func (r *repository) SelectByEmail(ctx context.Context, email string) (User, error) {
	var u User

	return u, r.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
}

func (r *repository) SelectByName(ctx context.Context, name string) (User, error) {
	var u User

	return u, r.db.WithContext(ctx).Where("name = ?", name).Select(&u).Error
}

func (r *repository) Create(ctx context.Context, user *User) error {
	return r.db.WithContext(ctx).Create(user).Error
}
