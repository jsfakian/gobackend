package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	u "gobackend/utils"
	"gobackend/db"
	"gobackend/models"
	"gobackend/configuration"
	"time"

	"github.com/gorilla/sessions"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")

func randSeq(n int) string {
	rand.Seed(time.Now().Unix())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// CreateUser creates an account for a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Create User")
	user := &db.User{}
	a := models.UserCtrls{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(http.StatusBadRequest, "Invalid request"))
		return
	}

	resp := a.Create(*user) //Create user
	u.Respond(w, resp)
}

// Authenticate authenticates users that try to login
func Authenticate(w http.ResponseWriter, r *http.Request) {

	store := sessions.NewCookieStore([]byte(randSeq(20)))
	account := &db.User{}
	a := models.UserCtrls{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(http.StatusBadRequest, "Invalid request"))
		return
	}

	session, err := store.Get(r, "session")
	if err != nil {
		u.Respond(w, u.Message(http.StatusInternalServerError, "Cannot get Session"))
		return
	}

	session.Options.MaxAge = int(configuration.Conf.ExpirationCookie)
	session.Save(r, w)

	resp := a.Login(account.Username, account.Password)
	u.Respond(w, resp)
}
