package handlers

import (
	"alef/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

func GetPageContent(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var contents []models.Content
	vars := mux.Vars(r)
	pageId, _ := vars["page-id"]
	result := db.Where("page_id = ?", pageId).Find(&contents)
	if result.Error != nil {
		respondInternalServerError(w)
	} else {
		respondJSON(w, http.StatusOK, contents)
	}
}
