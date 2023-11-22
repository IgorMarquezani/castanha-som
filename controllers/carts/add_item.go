package carts

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"castanha/database"
	"castanha/models/cart"
	"castanha/models/cart/item"
	"castanha/models/session"
	"castanha/models/user"
)

func AddItem(c echo.Context) error {
	var ctx = c.Request().Context()

	data := bytes.NewBuffer(make([]byte, 0, c.Request().ContentLength))
	io.Copy(data, c.Request().Body)

	cookie, err := c.Cookie("_Secure1")
	if err != nil {
		return c.String(http.StatusUnauthorized, "not loged in")
	}

	sessionRepo, err := session.NewRepository(database.GetDB())
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
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

	cartRepo, err := cart.NewRepository(database.GetDB())
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	Cart, err := cartRepo.SelectByUser(ctx, u.ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	cartItemRepo, err := item.NewRepository(database.GetDB())
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	Item := item.Item{
		CartID:      Cart.ID,
		ProductName: data.String(),
		Quantity:    1,
	}

	if err := cartItemRepo.Create(ctx, &Item); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.String(http.StatusAlreadyReported, "product already in cart")
		}
    if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return c.String(http.StatusBadRequest, "no such product")
    }

		return c.String(http.StatusInternalServerError, "internal server error")
	}

	return c.String(http.StatusCreated, "")
}
