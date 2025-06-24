package model

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Button struct {
	LunarModel
	Title string `gorm:"size:32" json:"title"`
	Name  string `gorm:"size:32" json:"name"`
	Group string `gorm:"size:32" json:"group"`
}

func (b Button) BeforeDelete(tx *gorm.DB) error {
	var list []RoleButton
	err := tx.Find(&list, "button_id = ?", b.ID).Delete(&list).Error
	logrus.Infof("删除按钮 %d ，关联删除角色按钮 %d个", b.ID, len(list))
	return err
}
