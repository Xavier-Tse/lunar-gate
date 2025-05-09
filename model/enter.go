package model

type Page struct {
	Limit int    `form:"limit"`
	Page  int    `form:"page"`
	Sort  string `form:"sort"`
	Key   string `form:"key"`
}
