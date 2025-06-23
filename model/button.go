package model

type Button struct {
	LunarModel
	Title string `gorm:"size:32" json:"title"`
	Name  string `gorm:"size:32" json:"name"`
	Group string `gorm:"size:32" json:"group"`
}
