package routes

import (
	"castanha/controllers/carts"
	"castanha/controllers/products"
	"castanha/controllers/render"
	"castanha/controllers/users"
	cmiddleware "castanha/middleware"
	"github.com/labstack/echo/v4"
)

func ServerSideRender(e *echo.Echo) {
	e.GET("/home", render.Home)
	e.GET("/products/upload", render.ProductForm)
	e.GET("/products/list", render.ListProducts)
	e.GET("/products/list/all", render.ListAllProducts)
	e.GET("/products/list/match", render.SearchProduct)
	e.GET("/products/info", render.ProductInfo)
	e.GET("/cart/view", render.Cart)
}

func UserRoutes(e *echo.Echo) {
	e.File("/", "./views/static/html/login.html")
	e.POST("/user/register", users.Register)
	e.POST("/user/login", users.Login)

	g := e.Group("/user/info/personal")
	g.GET("/my_name", users.OwnUserName)
}

func ProductRoutes(e *echo.Echo) {
	e.POST("/product/register", products.Register, cmiddleware.AuthAdmin)
}

func CartItemRoutes(e *echo.Echo) {
	e.POST("/cart/add_item", carts.AddItem)
}
