package render

import (
	"castanha/database"
	"castanha/models/session"
	"castanha/models/user"
	features "castanha/models/user_features"

	"net/http"

	"github.com/labstack/echo/v4"
)

func ProductForm(c echo.Context) error {
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

	featuresRepo, err := features.NewRepository(database.GetDB())
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	f, err := featuresRepo.SelectByUserID(ctx, u.ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	if !f.AdminAccess {
		return c.Redirect(http.StatusPermanentRedirect, "/static/html/401.html")
	}

	return c.Render(http.StatusOK, "productForm", nil)
}
