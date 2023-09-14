package main

import (
	"prakerja10/configs"
	"prakerja10/routes"

	"github.com/labstack/echo/v4"
)



func main(){
	configs.InitDatabase()
	e := echo.New()
	routes.InitRoute(e)
	e.Start(":8000")
}







