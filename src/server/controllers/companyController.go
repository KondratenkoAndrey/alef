package controllers

import (
	"encoding/json"
	"github.com/KondratenkoAndrey/alef/models"
	"net/http"
)

func GetCompanyInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pages := models.GetCompanyInfo()
	json.NewEncoder(w).Encode(pages)
}
