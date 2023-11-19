package render

import (
	"castanha/database"
	"castanha/models/product"
	"net/http"

	"github.com/labstack/echo/v4"
)

type List struct {
	Products []product.Product
	Len      int
}

func ListProducts(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		l   List
	)

	Type := c.QueryParam("type")

	productRepo, err := product.NewRepository(database.GetDB())
	if err != nil {
		c.File("./views/static/html/")
	}

	l.Products, err = productRepo.SelectByType(ctx, Type)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	l.Len = len(l.Products)

	return c.Render(http.StatusOK, "productList", &l)
}
