package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Title    string `gorm:"size:32,unique" json:"title"`
	UserList []User `gorm:"many2many:user_roles;joinForeignKey:RoleID;JoinReferences:UserID" json:"userList"`
	MenuList []Menu `gorm:"many2many:role_menus;joinForeignKey:RoleID;JoinReferences:MenuID" json:"menuList"`
}
