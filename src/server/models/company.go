package models

type CompanyInfo struct {
	Id        uint   `json:"-" gorm:"comment:Идентификатор компании"`
	ShortName string `json:"shortName" gorm:"size:64;comment:Краткое наименование"`
	FullName  string `json:"fullName" gorm:"size:256;comment:Полное наименование"`
	Phone     string `json:"phone" gorm:"size:18;comment:Телефон"`
	Email     string `json:"email" gorm:"size:64;comment:Электронная почта"'`
}

func (CompanyInfo) TableName() string {
	return "company_info"
}
