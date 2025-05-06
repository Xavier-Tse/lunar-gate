package model

import "gorm.io/gorm"

type RoleMenuModel struct {
	gorm.Model
	RoleID uint `json:"roleID"`
	Role   Role `gorm:"foreignKey:RoleID" json:"-"`
	MenuID uint `json:"menuID"`
	Menu   Menu `gorm:"foreignKey:MenuID" json:"-"`
}
