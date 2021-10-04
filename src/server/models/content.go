package models

type Content struct {
	Id     uint   `json:"id" gorm:"autoIncrement; comment:Идентификатор контента"`
	PageId string `json:"-" gorm:"index:idx_content_page_tag_unique; comment:Страница"`
	Tag    string `json:"tag" gorm:"size:100; index:idx_content_page_tag_unique; comment:Контент"`
	Data   string `json:"data" gorm:"comment:Тэг контента"`
	Page   Page   `json:"-"`
}

func (Content) TableName() string {
	return "content"
}
