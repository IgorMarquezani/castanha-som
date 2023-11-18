package products

import (
	"bytes"
	"castanha/database"
	"castanha/models/product"
	"castanha/models/product/description"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Register(c echo.Context) error {
	var (
		p   product.Product
		ctx = c.Request().Context()
	)

	JSON := c.FormValue("data")

	err := json.NewDecoder(bytes.NewBuffer([]byte(JSON))).Decode(&p)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid json format")
	}

	srcHeader, err := c.FormFile("image")
	if err != nil {
		return c.String(http.StatusBadRequest, "please send the product image")
	}

	str := strings.Split(srcHeader.Filename, ".")
	fileName := p.Name + "." + str[len(str)-1]

	file, err := os.OpenFile("./views/static/images/products/"+fileName, os.O_RDONLY, os.ModeSetuid)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	defer file.Close()

	if file != nil {
		return c.String(http.StatusAlreadyReported, "named product already exist")
	}

	src, err := srcHeader.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	defer src.Close()

	dst, err := os.Create("./views/static/images/products/" + fileName)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	defer dst.Close()

	io.Copy(dst, src)

	productRepo, err := product.NewRepository(database.GetDB())
	if err != nil {
		os.Remove("./views/static/images/products/" + fileName)
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	p.ImageName = fileName

	err = productRepo.Create(ctx, &p)
	if err != nil {
		os.Remove("./views/static/images/products/" + fileName)

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.String(http.StatusAlreadyReported, "named product already exist")
		}

		return c.String(http.StatusInternalServerError, "internal server error")
	}

	descriptionRepo, err := description.NewRepository(database.GetDB())
	if err != nil {
		os.Remove("./views/static/images/products/" + srcHeader.Filename)
		descriptionRepo.DeleteByProductName(ctx, p.Name)
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	for _, v := range p.Descriptions {
		v.ProductName = p.Name
		if err := descriptionRepo.Create(ctx, &v); err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				continue
			}

			os.Remove("./views/static/images/products/" + srcHeader.Filename)
			descriptionRepo.DeleteByProductName(ctx, p.Name)
			productRepo.Delete(ctx, p.Name)

			c.String(http.StatusInternalServerError, "internal server error")
			break
		}
	}

	return c.String(http.StatusCreated, "")
}
