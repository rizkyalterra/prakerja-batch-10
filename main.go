package main

import (
	"prakerja10/configs"
	"prakerja10/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)



func main(){
	loadEnv()
	configs.InitDatabase()
	e := echo.New()
	routes.InitRoute(e)
	e.Start(":8000")
}

func loadEnv(){
	err := godotenv.Load()
	if err != nil {
		panic("Failed load env file")
	}
}







