package models

type CompanyInfo struct {
	ShortName string `json:"shortName"`
	FullName  string `json:"fullName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

func (CompanyInfo) TableName() string {
	return "company_info"
}
