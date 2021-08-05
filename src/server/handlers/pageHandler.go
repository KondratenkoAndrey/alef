package handlers

import (
	"alef/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"net/url"
)

func GetPageByUrl(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var page models.Page
	vars := mux.Vars(r)
	pageUrl, _ := url.QueryUnescape(vars["url"])
	result := db.Where("url = ?", pageUrl).First(&page)
	if result.Error != nil {
		respondInternalServerError(w)
	} else {
		respondJSON(w, http.StatusOK, page)
	}
}
