package users

import (
	"castanha/database"
	"castanha/models/session"
	"castanha/models/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func OwnUserName(c echo.Context) error {
	var ctx = c.Request().Context()

	sessionRepo, err := session.NewRepository(database.GetDB())
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	cookie, err := c.Cookie("_Secure1")
	if err != nil {
		return c.String(http.StatusUnauthorized, "missing session cookie")
	}

	s, err := sessionRepo.SelectByKeyAccess(ctx, cookie.Value)
	if err != nil {
		return c.String(http.StatusUnauthorized, "invalid session cookie key accesss")
	}

	userRepo, err := user.NewRepository(database.GetDB())
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	u, err := userRepo.SelectByID(ctx, s.UserId)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	return c.String(http.StatusOK, u.Name)
}
