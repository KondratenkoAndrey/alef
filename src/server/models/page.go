package models

type Page struct {
	Id          uint      `json:"id" gorm:"autoIncrement; comment:Идентификатор страницы"`
	ParentId    string    `json:"-" gorm:"foreignKey:Page; index; comment:Идентификатор родительской страницы"`
	Url         string    `json:"url" gorm:"not null; size:255; index:idx_url_unique; comment:Относительный адрес"`
	Name        string    `json:"name" gorm:"size:50; comment:Имя страницы"`
	Title       string    `json:"title" gorm:"size:100; comment:Заголовок страницы"`
	Description string    `json:"description" gorm:"size: 255; comment:Краткое описание"`
	Pages       []Page    `json:"-" gorm:"foreignKey:ParentId"`
	Contents    []Content `json:"-" gorm:""`
}

func (Page) TableName() string {
	return "page"
}
