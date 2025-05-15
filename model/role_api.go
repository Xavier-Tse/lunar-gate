package model

import (
	"fmt"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RoleApi struct {
	LunarModel
	ApiID  uint `json:"apiID"`
	Api    Api  `gorm:"foreignKey:ApiID" json:"-"`
	RoleID uint `json:"roleID"`
	Role   Role `gorm:"foreignKey:RoleID" json:"role"`
}

func (a RoleApi) BeforeDelete(tx *gorm.DB) error {
	tx.Preload("Api").Take(&a)
	global.Casbin.RemovePolicy(fmt.Sprintf("%d", a.RoleID), a.Api.Path, a.Api.Method)
	logrus.Infof("删除角色api %d %s %s", a.RoleID, a.Api.Path, a.Api.Method)
	return nil
}
