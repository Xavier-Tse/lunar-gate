package model

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type User struct {
	LunarModel
	Username string `gorm:"size:32" json:"username"`
	Nickname string `gorm:"size:32" json:"nickname"`
	Avatar   string `gorm:"size:255" json:"avatar"`
	Email    string `gorm:"size:128" json:"email"`
	Password string `gorm:"size:64" json:"-"`
	IsAdmin  bool   `gorm:"default:false" json:"isAdmin"`
	RoleList []Role `gorm:"many2many:user_roles;joinForeignKey:UserID;JoinReferences:RoleID" json:"roleList"`
}

func (u User) BeforeDelete(tx *gorm.DB) error {
	var userRoleList []UserRole
	err := tx.Find(&userRoleList, "user_id = ?", u.ID).Delete(&userRoleList).Error
	logrus.Infof("删除用户角色 %d 条", len(userRoleList))
	return err
}

func (u User) GetRoleList() []uint {
	var roleList []uint
	for _, role := range u.RoleList {
		roleList = append(roleList, role.ID)
	}
	return roleList
}
