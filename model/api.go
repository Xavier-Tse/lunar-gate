package model

import "gorm.io/gorm"

type Api struct {
	gorm.Model
	Name   string `json:"name"`
	Path   string `json:"path"`
	Method string `json:"method"`
	Group  string `json:"group"`
}
