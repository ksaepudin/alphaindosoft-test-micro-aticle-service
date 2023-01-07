package entity

type Responses struct {
	Res Response `json:"response"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Pagination struct {
	Next      int `json:"next"`
	Prev      int `json:"prev"`
	TotalPage int `json:"total_page"`
}
