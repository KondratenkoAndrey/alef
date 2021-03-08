package main

import (
	"github.com/KondratenkoAndrey/alef/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/company-info", controllers.CompanyInfoHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
