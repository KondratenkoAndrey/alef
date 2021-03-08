package main

import (
	"github.com/KondratenkoAndrey/alef/controllers"
	"github.com/KondratenkoAndrey/alef/middlewares"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/company-info", controllers.GetCompanyInfo)
	env := os.Getenv("APP_ENV")
	if strings.ToLower(env) != "production" {
		router.Use(middlewares.LoggingMiddleware)
	}
	log.Fatal(http.ListenAndServe(":8080", router))
}
