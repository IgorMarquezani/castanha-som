package session

import (
	"context"
	"errors"
	"net/http"

	"castanha/database"
	"castanha/models/user"
	features "castanha/models/user_features"

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

func (a *authenticator) AuthAdminSession(c echo.Context) (int, error) {
	var ctx = c.Request().Context()

	cookie, err := c.Cookie("_Secure1")
	if err != nil {
		return http.StatusUnauthorized, errors.New("missing session cookie")
	}

	sessionRepo, err := NewRepository(a.DB)
	if err != nil {
		return http.StatusInternalServerError, errors.New("internal server error")
	}

	s, err := sessionRepo.SelectByKeyAccess(a.ctx, cookie.Value)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusUnauthorized, errors.New("invalid session cookie key accesss")
		}

		return http.StatusInternalServerError, errors.New("internal server error")
	}

	userRepo, err := user.NewRepository(database.GetDB())
	if err != nil {
		return http.StatusInternalServerError, err
	}

	u, err := userRepo.SelectByID(ctx, s.UserId)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	featuresRepo, err := features.NewRepository(database.GetDB())
	if err != nil {
		return http.StatusInternalServerError, err
	}

	f, err := featuresRepo.SelectByUserID(ctx, u.ID)
	if err != nil {
		return http.StatusInternalServerError, errors.New("not a admin")
	}

	if !f.AdminAccess {
		return http.StatusUnauthorized, err
	}

	return http.StatusOK, nil
}
