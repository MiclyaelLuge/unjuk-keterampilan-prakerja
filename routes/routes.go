package routes

import (
	"os"
	authcontroller "unjuk_prakerja/controllers/auth"
	productcontroller "unjuk_prakerja/controllers/products"
	usercontroller "unjuk_prakerja/controllers/user"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.POST("/login", authcontroller.LoginController)
	e.POST("/register", authcontroller.RegisterController)
	e.POST("/reg_products", productcontroller.RegisterProduct)

	eAuthUser := e.Group("")
	eAuthUser.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))
	eAuthUser.GET("/users", usercontroller.GetUsersController)
}
