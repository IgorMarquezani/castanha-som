package main

import (
	"castanha/routes"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var t = &Template{
	templates: template.Must(template.ParseGlob("./views/templates/*.html")),
}

func main() {
	e := echo.New()

	e.Renderer = t

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
	routes.ServerSideRender(e)
	routes.ProductRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
