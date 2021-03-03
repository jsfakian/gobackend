package models

import (
	"fmt"
	"net/http"
	"os"
	"gobackend/db"
	u "gobackend/utils"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

//Token is the JWT claims struct
type Token struct {
	UserID uint
	jwt.StandardClaims
}

//UserCtrls to control the account model
type UserCtrls struct{}

// Validate incoming user details...
func (a *UserCtrls) Validate(user db.User) (map[string]interface{}, bool) {

	fmt.Println("Validation")

	if len(user.Password) < 6 {
		return u.Message(http.StatusBadRequest, "Password is required"), false
	}

	//check for errors and duplicate username
	acc := db.GetUserByUsername(user.Username)
	if acc != nil {
		return u.Message(http.StatusNotAcceptable, "Username is already in use by another user."), false
	}

	return u.Message(http.StatusOK, "Requirement passed"), true
}

// Create function creates a new account
func (a *UserCtrls) Create(user db.User) map[string]interface{} {

	fmt.Println("Create")
	if _, ok := a.Validate(user); !ok {
		fmt.Println("Not validated")
		return nil
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	err := db.CreateObject(&user)

	if err != nil {
		return u.Message(http.StatusBadRequest, "Failed to create user, connection error.")
	}

	//Create new JWT token for the newly registered user
	tk := &db.Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString

	user.Password = "" //delete password

	response := u.Message(http.StatusCreated, "User has been created")
	response["user"] = user
	return response
}

// Login function
func (a *UserCtrls) Login(username, password string) map[string]interface{} {

	fmt.Println("Login")
	//check for errors and duplicate username
	user := db.GetUserByUsername(username)
	if user == nil {
		return u.Message(http.StatusUnauthorized, "Username not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(http.StatusUnauthorized, "Invalid login credentials. Please try again")
	}
	//Worked! Logged In
	user.Password = ""

	//Create JWT token
	tk := &Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString //Store the token in the response

	resp := u.Message(http.StatusOK, "Logged In")
	resp["token"] = user.Token
	return resp
}
