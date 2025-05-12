package model

type Api struct {
	LunarModel
	Name   string `gorm:"size:32" json:"name"`
	Path   string `gorm:"size:255" json:"path"`
	Method string `gorm:"size:16" json:"method"`
	Group  string `gorm:"size:32" json:"group"`
}
