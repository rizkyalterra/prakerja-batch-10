package controllers

import (
	"fmt"
	"net/http"
	"prakerja10/configs"
	userbase "prakerja10/models/base"
	userdatabase "prakerja10/models/user/database"
	userrequest "prakerja10/models/user/request"
	userresponse "prakerja10/models/user/response"

	"github.com/labstack/echo/v4"
)

func AddUsersController(c echo.Context) error {
	var userRegister userrequest.UserRegister
	c.Bind(&userRegister)
	fmt.Println("Daftar user userRegister", userRegister)

	if userRegister.Email == "" {
		return c.JSON(http.StatusBadRequest, userbase.BaseRespose{
			Status: false,
			Message: "Email still empty",
			Data: nil,
		})
	}

	var userDatabase userdatabase.User
	userDatabase.Name = userRegister.Name
	userDatabase.Email = userRegister.Email
	userDatabase.Password = userRegister.Password

	fmt.Println("Daftar user userDatabase", userDatabase)

	result := configs.DB.Create(&userDatabase)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, userbase.BaseRespose{
			Status: false,
			Message: "Failed add data users",
			Data: nil,
		})
	}

	var userResponse userresponse.UserResponse
	userResponse.MapFromDatabase(userDatabase)

	return c.JSON(http.StatusCreated, userbase.BaseRespose{
		Status: true,
		Message: "Success add data users",
		Data: userResponse,
	})
}

func LoginController(c echo.Context) error {
	var userLogin userrequest.UserLogin
	c.Bind(&userLogin)

	var userDatabase userdatabase.User
	userDatabase.MapFromLogin(userLogin)

	if userDatabase.Email == "alterra@gmail.com" && userDatabase.Password == "123ABC" {
		return c.JSON(http.StatusOK, userbase.BaseRespose{
			Status: true,
			Message: "Success login",
			Data: userLogin,
		})
	} 
	return c.JSON(http.StatusUnauthorized, userbase.BaseRespose{
		Status: false,
		Message: "Failed login",
		Data: nil,
	})
}

func GetUsersController(c echo.Context) error {
	var users []userdatabase.User

	result := configs.DB.Find(&users)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, userbase.BaseRespose{
			Status: false,
			Message: "Failed get data users",
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, userbase.BaseRespose{
		Status: true,
		Message: "Success get data users",
		Data: users,
	})
}
