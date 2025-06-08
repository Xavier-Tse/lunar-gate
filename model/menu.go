package model

import (
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Menu struct {
	LunarModel
	Key             uint    `gorm:"-" json:"key"`
	Enable          bool    `gorm:"default:false" json:"enable"`
	Name            string  `gorm:"size:32,unique" json:"name"`
	Path            string  `gorm:"size:64" json:"path"`
	Component       string  `gorm:"size:128" json:"component"`
	Meta            Meta    `gorm:"embedded" json:"meta"`
	ParentMenuID    *uint   `json:"parentMenuID"`
	ParentMenuModel *Menu   `gorm:"foreignKey:ParentMenuID" json:"-"`
	Children        []*Menu `gorm:"foreignKey:ParentMenuID" json:"children,omitempty"`
	Sort            int     `json:"sort"`
}

type Meta struct {
	Icon  string `gorm:"size:255" json:"icon"`
	Title string `gorm:"size:16" json:"title"`
}

func (u Menu) BeforeDelete(tx *gorm.DB) error {
	tx.Preload("Children").Take(&u)
	for _, child := range u.Children {
		tx.Model(child).Update("parent_menu_id", u.ParentMenuID)
		if u.ParentMenuID == nil {
			logrus.Infof("删除菜单 %d,修改菜单 %d 的parent_menu_id为 null", u.ID, child.ID)
		} else {
			logrus.Infof("删除菜单 %d,修改菜单 %d 的parent_menu_id为 %d", u.ID, child.ID, *u.ParentMenuID)
		}
	}

	var roleMenuList []RoleMenu
	err := tx.Find(&roleMenuList, "menu_id = ?", u.ID).Delete(&roleMenuList).Error
	logrus.Infof("删除菜单 %d,删除角色菜单 %d 条", u.ID, len(roleMenuList))
	return err
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

func SubMenuList(model *Menu) {
	global.DB.Preload("Children").Take(&model)
	model.Key = model.ID
	for _, child := range model.Children {
		SubMenuList(child)
	}
}
