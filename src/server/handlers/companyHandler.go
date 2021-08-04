package handlers

import (
	"alef/models"
	"gorm.io/gorm"
	"net/http"
)

func GetCompanyInfo(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	companyInfo, err := models.GetCompanyInfo(db)
	if err != nil {
		respondError(w, 500, "Внутренняя ошибка сервера. Повторите попытку позже.")
	} else {
		respondJSON(w, http.StatusOK, companyInfo)
	}
}
