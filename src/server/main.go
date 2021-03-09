package main

import (
	"alef/controllers"
	"alef/middlewares"
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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server started with environment:", env)
	log.Println("Listen on", "localhost:"+port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
