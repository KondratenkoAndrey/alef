package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func CompanyInfoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	w.Header().Set("Content-Type", "application/json")
	pages := getCompanyInfo()
	json.NewEncoder(w).Encode(pages)
}

func getCompanyInfo() CompanyInfo {
	return CompanyInfo{
		ShortName: "Алеф",
		FullName:  "Компания Алеф",
		Phone:     "+7(789)123-45-67",
		Email:     "info@example.ru",
	}
}

type CompanyInfo struct {
	ShortName string `json:"shortName"`
	FullName  string `json:"fullName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}
