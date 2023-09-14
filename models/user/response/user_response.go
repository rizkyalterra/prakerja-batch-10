package response

import (
	"prakerja10/middlewares"
	"prakerja10/models/user/database"
)

type UserResponse struct {
	Id int 			`json:"id" gorm:"primaryKey autoIncrement"`
	Name string 	`json:"name"`
	Email string 	`json:"email"`
	Token string 	`json:"token"`
}

func (userResponse *UserResponse) MapFromDatabase(userDatabase database.User) {
	userResponse.Id = userDatabase.Id
	userResponse.Name = userDatabase.Name
	userResponse.Email = userDatabase.Email
	userResponse.Token = middlewares.GenerateJWT(userDatabase)
}