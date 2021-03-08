package controllers

import (
	"encoding/json"
	"github.com/KondratenkoAndrey/alef/models"
	"log"
	"net/http"
)

func GetCompanyInfo(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	w.Header().Set("Content-Type", "application/json")
	pages := models.GetCompanyInfo()
	json.NewEncoder(w).Encode(pages)
}
