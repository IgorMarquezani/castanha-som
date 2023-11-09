package routes

import (
	"castanha/controllers/users"
	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	e.File("/", "./views/static/html/login.html")
	e.POST("/user/register", users.Register)
	e.POST("/user/login", users.Login)
}
