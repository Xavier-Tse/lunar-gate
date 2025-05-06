package model

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	UserID uint `json:"userID"`
	User   User `gorm:"foreignKey:UserID" json:"-"`
	RoleID uint `json:"roleID"`
	Role   Role `gorm:"foreignKey:RoleID" json:"-"`
}
