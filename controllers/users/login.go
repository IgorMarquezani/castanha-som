package users

import (
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"castanha/database"
	"castanha/hasher"
	"castanha/models/session"
	"castanha/models/user"
)

func Login(c echo.Context) error {
	var (
		email  = c.FormValue("email")
		passwd = c.FormValue("password")
		ctx    = c.Request().Context()
		now    = time.Now()
	)

	userRepo, err := user.NewRepository(database.GetDB())
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	u, err := userRepo.SelectByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.String(http.StatusUnauthorized, "invalid credentials")
		}

		return c.String(http.StatusInternalServerError, "internal server error")
	}

	if err := hasher.CompareHashAndPassword(u.Passwd, passwd); err != nil {
		return c.String(http.StatusUnauthorized, "invalid credentials")
	}

	sessionRepo, err := session.NewRepository(database.GetDB())
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	s := session.Session{
		KeyAccess: uuid.NewString(),
		UserId:    u.ID,
		ExpiresAt: now.Add(time.Hour * 24).Format("2006-01-02 15:04:05"),
	}

	err = sessionRepo.Create(ctx, s)
	if err != nil && errors.Is(err, gorm.ErrDuplicatedKey) {
		oldSession, err := sessionRepo.SelectByUserId(ctx, u.ID)
		if err != nil {
			return c.String(http.StatusInternalServerError, "internal server error")
		}
		s.KeyAccess = oldSession.KeyAccess

		err = sessionRepo.UpdateDuration(ctx, u.ID, now.Add(time.Hour*24).Format("2006-01-02 15:04:05"))
		if err != nil {
			return c.String(http.StatusInternalServerError, "internal server error")
		}
	} else if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	cookie := http.Cookie{
		Name:    "_Secure1",
		Value:   s.KeyAccess,
		Expires: now.Add(time.Hour * 24),
	}

	c.SetCookie(&cookie)

	return c.String(http.StatusOK, "")
}
