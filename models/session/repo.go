package session

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

func (r *repository) Create(ctx context.Context, session Session) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	return r.db.WithContext(ctx).Create(session).Error
}

func (r *repository) SelectByUserId(ctx context.Context, userId string) (Session, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	session := Session{}

	err := r.db.WithContext(ctx).Where("user_id = ?", userId).First(&session).Error

	return session, err
}

func (r *repository) SelectByKeyAccess(ctx context.Context, sessionId string) (Session, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	session := Session{}

	return session, r.db.WithContext(ctx).Where("key_access = ?", sessionId).First(&session).Error
}

func (r *repository) UpdateDuration(ctx context.Context, userId, newDuration string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	operation := r.db.Model(&Session{}).WithContext(ctx).Where("user_id = ?", userId)

	return operation.Update("expires_at", newDuration).Error
}
