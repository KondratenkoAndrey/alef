package models

type CompanyInfo struct {
	ShortName string `json:"shortName"`
	FullName  string `json:"fullName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

func GetCompanyInfo() CompanyInfo {
	return CompanyInfo{
		ShortName: "Алеф",
		FullName:  "Компания Алеф",
		Phone:     "+7(789)123-45-67",
		Email:     "info@example.ru",
	}
}
