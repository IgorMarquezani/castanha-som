package users

import (
	"errors"
	"net/http"
	"net/mail"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"castanha/database"
	"castanha/hasher"
	"castanha/models/user"
)

func Register(c echo.Context) error {
	ctx := c.Request().Context()
	name := c.FormValue("name")
	email := c.FormValue("email")
	passwd := c.FormValue("password")

	if len(name) == 0 {
		return c.String(http.StatusBadRequest, "name cannot be empty")
	}

	e, err := mail.ParseAddress(email)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid email format")
	}
	email = e.Address

	validator := user.NewValidator(name, passwd)

	if err := validator.ValidateName(); err != nil {
		c.Response().Header().Set("Content-Type", "application/json; charset=utf-8")
		return c.JSONPretty(http.StatusBadRequest, err, "")
	}

	if err := validator.ValidatePassword(); err != nil {
		c.Response().Header().Set("Content-Type", "application/json; charset=utf-8")
		return c.JSONPretty(http.StatusBadRequest, err, "")
	}

	repo, err := user.NewRepository(database.GetDB())
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	hash, err := hasher.HashPassword(passwd)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	u := user.User{
		ID:     uuid.New().String(),
		Name:   name,
		Email:  email,
		Passwd: hash,
	}

	if err := repo.Create(ctx, &u); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.String(http.StatusAlreadyReported, "email already in use")
		}

		return c.String(http.StatusInternalServerError, "internal server error")
	}

	return c.String(http.StatusCreated, "user created")
}
