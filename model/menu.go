package model

import "github.com/lunarise-dev/lunar-gate/global"

type Menu struct {
	LunarModel
	Enable          bool    `gorm:"default:false" json:"enable"`
	Name            string  `gorm:"size:32,unique" json:"name"`
	Path            string  `gorm:"size:64" json:"path"`
	Component       string  `gorm:"size:128" json:"component"`
	Meta            Meta    `gorm:"embedded" json:"meta"`
	ParentMenuID    *uint   `json:"parentMenuID"`
	ParentMenuModel *Menu   `gorm:"foreignKey:ParentMenuID" json:"-"`
	Children        []*Menu `gorm:"foreignKey:ParentMenuID" json:"children"`
	Sort            int     `json:"sort"`
}

type Meta struct {
	Icon  string `gorm:"size:255" json:"icon"`
	Title string `gorm:"size:16" json:"title"`
}

func FindSubMenuList(model Menu) []Menu {
	global.DB.Preload("Children").Take(&model)

	var list []Menu
	list = append(list, model)
	for _, child := range model.Children {
		list = append(list, FindSubMenuList(*child)...)
	}
	return list
}
