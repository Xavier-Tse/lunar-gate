package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:32" json:"username"`
	Nickname string `gorm:"size:32" json:"nickname"`
	Avatar   string `gorm:"size:255" json:"avatar"`
	Email    string `gorm:"size:128" json:"email"`
	Password string `gorm:"size:64" json:"-"`
	IsAdmin  bool   `gorm:"default:false" json:"isAdmin"`
	RoleList []Role `gorm:"many2many:user_roles;joinForeignKey:UserID;JoinReferences:RoleID" json:"roleList"`
}
