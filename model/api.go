package model

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Api struct {
	LunarModel
	Name   string `gorm:"size:32" json:"name"`
	Path   string `gorm:"size:255" json:"path"`
	Method string `gorm:"size:16" json:"method"`
	Group  string `gorm:"size:32" json:"group"`
}

func (a Api) BeforeDelete(tx *gorm.DB) error {
	var list []RoleApi
	err := tx.Find(&list, "api_id = ?", a.ID).Delete(&list).Error
	logrus.Infof("删除api %d ，关联删除角色api %d 个", a.ID, len(list))
	return err
}
