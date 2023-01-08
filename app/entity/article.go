package entity

type ArticlesListRequest struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
type SearchArticlesRequest struct {
	Search string `json:"search" query:"search"`
	Author string `json:"author" gorm:"column:author" query:"author"`
}

type ArticlesListResponse struct {
	Articles []*Articles
	Page     *Pagination
}

type Articles struct {
	Id        string `json:"id" gorm:"column:id" query:"id"`
	Author    string `json:"author" gorm:"column:author" query:"author"`
	Title     string `json:"title" gorm:"column:title" query:"title"`
	Body      string `json:"body" gorm:"column:body" query:"body"`
	CreatedAt string `json:"created" gorm:"column:created" query:"created"`
}

func (Articles) TableName() string {
	return "tbl_articles"
}
