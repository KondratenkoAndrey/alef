package models

type Service struct {
	Id          uint      `json:"-" gorm:"comment:Идентификатор услуги"`
	Name        string    `json:"name" gorm:"size:255; unique index; comment:Название услуги"`
	Description string    `json:"description" gorm:"size:255; comment:Краткое описание"`
	Content     string    `json:"content" gorm:"comment:Содежание"`
	Tag         string    `json:"tag" gorm:"size:50; not null; unique index; comment:Тэг услуги"`
	ImagePath   string    `json:"imagePath" gorm:"size:255; comment:Изображение"`
	Order       uint      `json:"order" gorm:"not null; unique index; comment:Порядок отображения"`
	Caption     string    `json:"caption" gorm:"foreignKey:Service; index; comment:Заголовок услуги"`
	Title       string    `json:"title" gorm:"comment:Заголовок страницы услуги"`
	ParentId    uint      `json:"parentId" gorm:"comment:Идентификатор родительской услуги"`
	Services    []Service `json:"-" gorm:"foreignKey:ParentId; comment:"`
}

func (Service) TableName() string {
	return "service"
}
