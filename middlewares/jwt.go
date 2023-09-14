package middlewares

import (
	"prakerja10/models/user/database"
	"time"

	"github.com/golang-jwt/jwt/v4"
)


type jwtCustomClaims struct {
	UserId  int `json:"userId"`
	Name  string `json:"name"`
	jwt.RegisteredClaims
}

func GenerateJWT(user database.User) string {
	claims := &jwtCustomClaims{
		user.Id,
		user.Name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, _ := token.SignedString([]byte("123"))
	
	return t
}