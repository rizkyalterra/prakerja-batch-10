package routes

import (
	"prakerja10/controllers"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.POST("/login", controllers.LoginController)

	eAuth := e.Group("")
	eAuth.Use(echojwt.JWT([]byte("123")))
	eAuth.GET("/users", controllers.GetUsersController)
	eAuth.POST("/users", controllers.AddUsersController)
	
}