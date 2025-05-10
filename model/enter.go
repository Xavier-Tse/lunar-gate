package model

type Page struct {
	Limit int    `form:"limit"`
	Page  int    `form:"page"`
	Sort  string `form:"sort"`
	Key   string `form:"key"`
}

type IDRequest struct {
	ID uint `json:"id" form:"id" url:"id"`
}

type IDListRequest struct {
	IDList []uint `json:"idList"`
}
