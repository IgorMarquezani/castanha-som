package render

import (
	"net/http"

	"castanha/database"
	"castanha/models/product"

	"github.com/labstack/echo/v4"
)

func ListAllProducts(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		l   productList
	)

	productRepo, err := product.NewRepository(database.GetDB())
	if err != nil {
		c.File("./views/static/html/")
	}

	l.Products, err = productRepo.Select(ctx)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	l.Len = len(l.Products)

	return c.Render(http.StatusOK, "productList", &l)
}
