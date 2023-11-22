package render

import "github.com/labstack/echo/v4"

func Cart(c echo.Context) error {
  return c.File("./views/static/html/cart.html")
}
