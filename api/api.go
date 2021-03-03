package api

import (

	"github.com/gorilla/mux"

	"gobackend/controllers"
	"gobackend/app"
)

//InitRoutes initializes the available routes
func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	// User endpoints
	router.HandleFunc("/user/create", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/login", controllers.Authenticate).Methods("POST")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	return router
}
