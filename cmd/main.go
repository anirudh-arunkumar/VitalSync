package main

import (
	"log"
	"net/http"
	"VitalSync/config"
	"VitalSync/routes"
)

func main() {
	config.ConnectDB()
	router := routes.SetupRoutes()
	log.Fatal(http.ListenAndServe(":8000", router))
}