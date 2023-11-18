package middleware

import (
	"castanha/database"
	"castanha/models/session"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var ctx = c.Request().Context()

		auth, err := session.NewAuthenticator(ctx, database.GetDB())
		if err != nil {
			return c.String(http.StatusInternalServerError, "internal server error")
		}

		status, err := auth.AuthSession(c)
		if err != nil {
			return c.String(status, err.Error())
		}

		return next(c)
	}
}

func AuthAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var ctx = c.Request().Context()

		auth, err := session.NewAuthenticator(ctx, database.GetDB())
		if err != nil {
			return c.String(http.StatusInternalServerError, "internal server error")
		}

		status, err := auth.AuthAdminSession(c)
		if status == http.StatusUnauthorized {
			return c.Redirect(http.StatusPermanentRedirect, "/static/html/401.html")
		}

		if status != http.StatusOK || err != nil {
			return c.String(status, err.Error())
		}

		return next(c)
	}
}
