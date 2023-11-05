package main

import (
	"castanha/routes"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.Static("/static/", "./views/static/")

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError

		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}

		if code == http.StatusNotFound {
			c.File("./views/static/html/404.html")
		} else {
			c.File("./views/static/html/500.html")
		}
	}

	routes.UserRoutes(e)

	e.File("/", "./views/static/html/login.html")

	e.Logger.Fatal(e.Start(":8080"))
}
