package render

import (
	"castanha/database"
	"castanha/models/product"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func SearchProduct(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		l   productList
	)

	match := c.QueryParam("m")
	if match == "" {
		return c.String(http.StatusBadRequest, "no match sent")
	}

	query := strings.Replace(fullProductMatch, "?", match, 3)
	println(query)

	productRepo, err := product.NewRepository(database.GetDB())
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	l.Products, err = productRepo.RawSelect(ctx, query)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	l.Len = len(l.Products)

	return c.Render(http.StatusOK, "productList", &l)
}
