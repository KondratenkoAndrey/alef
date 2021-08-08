package models

type Content struct {
	Id     string `json:"-"`
	PageId string `json:"-"`
	Tag    string `json:"tag"`
	Data   string `json:"data"`
}

func (Content) TableName() string {
	return "content"
}
