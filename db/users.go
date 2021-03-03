package db

import (
	"github.com/dgrijalva/jwt-go"
)

//Token is the JWT claims struct
type Token struct {
	UserID uint
	jwt.StandardClaims
}

// User struct is a struct to rep user account
type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

//GetUser function
func GetUser(u uint) *User {
	acc := &User{}
	GetDB().Table("users").Where("id = ?", u).First(acc)
	if acc.Username == "" { //User not found!
		return nil
	}

	acc.Password = ""

	return acc
}

//GetUserByUsername function
func GetUserByUsername(username string) *User {
	acc := &User{}
	err := GetDB().Table("users").Where("username = ?", username).First(acc).Error

	if err != nil || acc.Username == "" { //User not found!
		return nil
	}

	acc.Password = ""

	return acc
}
