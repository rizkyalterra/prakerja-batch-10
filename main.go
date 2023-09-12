package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id int 			`json:"id"`
	Name string 	`json:"name"`
	Email string 	`json:"email"`
}

type BaseRespose struct {
	Status bool 		`json:"status"`
	Message string 		`json:"message"`
	Data interface{} 	`json:"data"`
}

func main(){
	e := echo.New()
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserDetailController)
	e.Start(":8000")
}


func GetUserDetailController(c echo.Context) error {
	id := c.Param("id")

	var users User = User{1, id, "A"}

	return c.JSON(http.StatusOK, BaseRespose{
		Status: true,
		Message: "Success get detail users",
		Data: users,
	})
}

func GetUsersController(c echo.Context) error {
	var users []User

	users = append(users, User{1, "A", "A"})
	users = append(users, User{2, "B", "B"})
	users = append(users, User{3, "C", "C"})

	return c.JSON(http.StatusOK, BaseRespose{
		Status: true,
		Message: "Success get data users",
		Data: users,
	})
}





