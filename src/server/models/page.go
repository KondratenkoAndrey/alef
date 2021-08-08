package models

type Page struct {
	Id          string `json:"id"`
	ParentId    string `json:"-"`
	Url         string `json:"url"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (Page) TableName() string {
	return "page"
}
