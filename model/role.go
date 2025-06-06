package model

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Role struct {
	LunarModel
	Title    string `gorm:"size:32,unique" json:"title"`
	UserList []User `gorm:"many2many:user_roles;joinForeignKey:RoleID;JoinReferences:UserID" json:"userList"`
	MenuList []Menu `gorm:"many2many:role_menus;joinForeignKey:RoleID;JoinReferences:MenuID" json:"menuList"`
	ApiList  []Api  `gorm:"many2many:role_apis;joinForeignKey:RoleID;JoinReferences:ApiID" json:"apiList"`
}

func (r Role) BeforeDelete(tx *gorm.DB) error {
	var roleMenuList []RoleMenu
	err := tx.Find(&roleMenuList, "role_id = ?", r.ID).Delete(&roleMenuList).Error
	logrus.Infof("删除角色菜单 %d 条", len(roleMenuList))

	var roleUserList []UserRole
	err = tx.Find(&roleUserList, "role_id = ?", r.ID).Delete(&roleUserList).Error
	logrus.Infof("删除角色用户 %d 条", len(roleUserList))

	var roleApiList []UserRole
	err = tx.Find(&roleApiList, "role_id = ?", r.ID).Delete(&roleApiList).Error
	logrus.Infof("删除角色Api %d 条", len(roleApiList))
	return err
}
