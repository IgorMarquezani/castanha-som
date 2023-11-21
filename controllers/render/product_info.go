package render

import (
	"castanha/database"
	"castanha/models/product"
	"castanha/models/product/description"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ProductInfo(c echo.Context) error {
  var (
    ctx = c.Request().Context()
  )

  name := c.QueryParam("name")
  if len(name) < 1 {
    return c.String(http.StatusBadRequest, "please give a name")
  }

	productRepo, err := product.NewRepository(database.GetDB())
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

  p, err := productRepo.FirstByName(ctx, name)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

  descRepo, err := description.NewRepository(database.GetDB())
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

  p.Descriptions, err = descRepo.SelectByProduct(ctx, p.Name)
  if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
  }

  return c.Render(http.StatusOK, "productPage", &p)
}
