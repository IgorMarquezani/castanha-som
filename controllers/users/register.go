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
	features "castanha/models/user_features"
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

	userRepo, err := user.NewRepository(database.GetDB())
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

	if err := userRepo.Create(ctx, &u); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.String(http.StatusAlreadyReported, "email already in use")
		}

		return c.String(http.StatusInternalServerError, "internal server error")
	}

  featuresRepo, err := features.NewRepository(database.GetDB())
  if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
  }

  userFeatures := features.UserFeatures {
    AdminAccess: false,
    UserID: u.ID,
  }

  if err := featuresRepo.Create(ctx, &userFeatures); err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
  }

	return c.String(http.StatusCreated, "user created")
}
