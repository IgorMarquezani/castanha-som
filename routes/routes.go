package routes

import (
  "github.com/labstack/echo/v4"
  "castanha/controllers/users"
)

func UserRoutes(e *echo.Echo) {
  e.POST("/user/register", users.Register)
  e.POST("/user/login", users.Login)
}
