package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id int 			`json:"id" gorm:"primaryKey autoIncrement"`
	Name string 	`json:"name"`
	Email string 	`json:"email"`
}

type UserLogin struct {
	Email string 		`json:"email"`
	Password string 	`json:"password"`
}


type BaseRespose struct {
	Status bool 		`json:"status"`
	Message string 		`json:"message"`
	Data interface{} 	`json:"data"`
}

func main(){
	InitDatabase()
	e := echo.New()
	e.GET("/users", GetUsersController)
	e.POST("/users", AddUsersController)
	e.GET("/users/:id", GetUserDetailController)
	e.POST("/login", LoginController)
	e.Start(":8000")
}


var DB *gorm.DB

func InitDatabase(){
	dsn := "root:123ABC4d.@tcp(127.0.0.1:3306)/prakerja10?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("gagal init database")
	}
	Migration()
}

func Migration(){
	DB.AutoMigrate(&User{})
}

func AddUsersController(c echo.Context) error {
	var user User
	c.Bind(&user)

	result := DB.Create(&user)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, BaseRespose{
			Status: false,
			Message: "Failed add data users",
			Data: nil,
		})
	}

	return c.JSON(http.StatusCreated, BaseRespose{
		Status: true,
		Message: "Success add data users",
		Data: user,
	})
}

func LoginController(c echo.Context) error {
	var userLogin UserLogin
	c.Bind(&userLogin)

	if userLogin.Email == "alterra@gmail.com" && userLogin.Password == "123ABC" {
		return c.JSON(http.StatusOK, BaseRespose{
			Status: true,
			Message: "Success login",
			Data: userLogin,
		})
	} 
	return c.JSON(http.StatusUnauthorized, BaseRespose{
		Status: false,
		Message: "Failed login",
		Data: nil,
	})
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

	result := DB.Find(&users)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, BaseRespose{
			Status: false,
			Message: "Failed get data users",
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, BaseRespose{
		Status: true,
		Message: "Success get data users",
		Data: users,
	})
}





