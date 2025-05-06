package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `size:"16" json:"username"`
	Nickname string `size:"32" json:"nickname"`
	Avatar   string `size:"255" json:"avatar"`
	Email    string `size:"128" json:"email"`
	Password string `size:"64" json:"-"`
	RoleList []Role `gorm:"many2many:user_roles;joinForeignKey:UserID;JoinReferences:RoleID" json:"roleList"`
}
