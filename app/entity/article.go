package entity

type ArticlesListRequest struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type ArticlesListResponse struct {
	Articles []*Articles
	Page     *Pagination
}

type Articles struct {
	Id        string `json:"id" gorm:"column:id" `
	Author    string `json:"author" gorm:"column:author" `
	Title     string `json:"title" gorm:"column:title" `
	Body      string `json:"body" gorm:"column:body" `
	CreatedAt string `json:"created" gorm:"column:created" `
}

func (Articles) TableName() string {
	return "tbl_articles"
}
