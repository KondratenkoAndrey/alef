package models

import (
	"gorm.io/gorm"
)

type CompanyInfo struct {
	ShortName string `json:"shortName"`
	FullName  string `json:"fullName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

func (CompanyInfo) TableName() string {
	return "company_info"
}

func GetCompanyInfo(db *gorm.DB) (CompanyInfo, error) {
	var companyInfo CompanyInfo
	db.Take(&companyInfo)

	if result := db.Take(&companyInfo); result.Error != nil {
		return companyInfo, result.Error
	}

	return companyInfo, nil
}
