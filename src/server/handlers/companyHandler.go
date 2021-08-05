package handlers

import (
	"alef/models"
	"gorm.io/gorm"
	"net/http"
)

func GetCompanyInfo(db *gorm.DB, w http.ResponseWriter) {
	var companyInfo models.CompanyInfo
	if result := db.Take(&companyInfo); result.Error != nil {
		respondInternalServerError(w)
	} else {
		respondJSON(w, http.StatusOK, companyInfo)
	}
}
