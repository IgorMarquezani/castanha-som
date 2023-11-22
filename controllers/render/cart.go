package render

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"castanha/database"
	"castanha/models/cart"
	"castanha/models/cart/item"
	"castanha/models/product"
	"castanha/models/session"
	"castanha/models/user"
)

type Purchase struct {
	Items []CartItem
	Total float32
}

type CartItem struct {
	Name      string
	Quantity  int
	Value     float64
	ImageName string
}

func Cart(c echo.Context) error {
	var ctx = c.Request().Context()
	var purchase Purchase

	cookie, err := c.Cookie("_Secure1")
	if err != nil {
		return c.String(http.StatusUnauthorized, "not loged in")
	}

	sessionRepo, err := session.NewRepository(database.GetDB())
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	s, err := sessionRepo.SelectByKeyAccess(ctx, cookie.Value)
	if err != nil {
		return c.String(http.StatusUnauthorized, "invalid session cookie key accesss")
	}

	userRepo, err := user.NewRepository(database.GetDB())
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	u, err := userRepo.SelectByID(ctx, s.UserId)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	cartRepo, err := cart.NewRepository(database.GetDB())
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	Cart, err := cartRepo.SelectByUser(ctx, u.ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	cartItemRepo, err := item.NewRepository(database.GetDB())
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	Items, err := cartItemRepo.SelectByCarID(ctx, Cart.ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	productRepo, err := product.NewRepository(database.GetDB())
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	purchase.Items = make([]CartItem, len(Items))

	for i, v := range Items {
		p, err := productRepo.FirstByName(ctx, v.ProductName)
		if err != nil {

			return c.String(http.StatusInternalServerError, "internal server error")
		}

		purchase.Total += float32(v.Quantity) * p.InCashValue

		purchase.Items[i] = CartItem{
			Name:      v.ProductName,
			Quantity:  v.Quantity,
			Value:     float64(p.InCashValue),
			ImageName: p.ImageName,
		}
	}

	return c.Render(http.StatusOK, "cart", &purchase)
}
