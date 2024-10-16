package routes

import (
	"github.com/gorilla/mux"
	"VitalSync/controllers"
)

func setupEndpoints() *mux.Router {
	router := mux.NewRouter()

	router.HandlerFunc("/register", controllers.registerUser).Methods("POST")
	router.HandlerFunc("/login", controllers.loginUser).Methods("POST")
	router.HandlerFunc("/fitness", controllers.addData).Methods("POST")
	router.HandlerFunc("/fitness/{userId}", controllers.getData).Methods("GET")

	return router
}