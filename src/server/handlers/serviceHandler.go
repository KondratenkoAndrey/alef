package handlers

import (
	"alef/models"
	"gorm.io/gorm"
	"net/http"
)

func GetAllServices(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var services []models.Service
	result := db.Find(&services)
	if result.Error != nil {
		respondInternalServerError(w)
	} else {
		respondJSON(w, http.StatusOK, services)
	}
}
