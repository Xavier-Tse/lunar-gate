package flags

import (
	"github.com/lunarise-dev/lunar-gate/global"
	"github.com/lunarise-dev/lunar-gate/model"
	"github.com/sirupsen/logrus"
)

func AutoMigrate() {
	if err := global.DB.AutoMigrate(
		&model.Api{},
		&model.Menu{},
		&model.Role{},
		&model.RoleMenu{},
		&model.User{},
		&model.UserRole{},
	); err != nil {
		logrus.Fatalf("表结构迁移失败 %v", err)
	}
	logrus.Infof("表结构迁移成功")
}
