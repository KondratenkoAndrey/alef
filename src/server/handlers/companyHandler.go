package controllers

import (
	"alef/models"
	"encoding/json"
	"net/http"
)

func GetCompanyInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pages := models.GetCompanyInfo()
	json.NewEncoder(w).Encode(pages)
}
