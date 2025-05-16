package model

import (
	"fmt"
	"github.com/Xavier-Tse/lunar-gate/global"
	"gorm.io/gorm"
)

type UserRole struct {
	LunarModel
	UserID uint `json:"userID"`
	User   User `gorm:"foreignKey:UserID" json:"-"`
	RoleID uint `json:"roleID"`
	Role   Role `gorm:"foreignKey:RoleID" json:"-"`
}

func (u UserRole) AfterDelete(tx *gorm.DB) error {
	global.Casbin.RemoveGroupingPolicy(fmt.Sprintf("%d", u.UserID), fmt.Sprintf("%d", u.RoleID))
	return nil
}

func (u UserRole) AfterCreate(tx *gorm.DB) error {
	global.Casbin.AddRoleForUser(fmt.Sprintf("%d", u.UserID), fmt.Sprintf("%d", u.RoleID))
	return nil
}
