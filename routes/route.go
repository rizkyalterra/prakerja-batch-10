package routes

import (
	"prakerja10/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	
	e.GET("/users", controllers.GetUsersController)
	e.POST("/users", controllers.AddUsersController)
	e.POST("/login", controllers.LoginController)
}