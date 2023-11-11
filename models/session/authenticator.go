package session

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type authenticator struct {
	ctx context.Context
	DB  *gorm.DB
}

func NewAuthenticator(ctx context.Context, db *gorm.DB) (*authenticator, error) {
	if db == nil {
		return nil, errors.New("database connection cannot be nil")
	}

	return &authenticator{
		ctx: ctx,
		DB:  db,
	}, nil
}

func (a *authenticator) AuthSession(c echo.Context) (int, error) {
	cookie, err := c.Cookie("_Secure1")
	if err != nil {
		return http.StatusUnauthorized, errors.New("missing session cookie")
	}

	repo, err := NewRepository(a.DB)
	if err != nil {
		return http.StatusInternalServerError, errors.New("internal server error")
	}

	_, err = repo.SelectByKeyAccess(a.ctx, cookie.Value)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusUnauthorized, errors.New("invalid session cookie key accesss")
		}

		return http.StatusInternalServerError, errors.New("internal server error")
	}

	return http.StatusOK, nil
}
