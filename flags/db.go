package flags

import (
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/lunarise-dev/lunar-gate/global"
	"github.com/lunarise-dev/lunar-gate/model"
	"github.com/sirupsen/logrus"
)

func AutoMigrate() {
	if err := global.DB.AutoMigrate(
		&model.Api{},
		&model.Menu{},
		&model.Role{},
		&model.RoleApi{},
		&model.RoleMenu{},
		&model.User{},
		&model.UserRole{},
		&gormadapter.CasbinRule{},
	); err != nil {
		logrus.Fatalf("表结构迁移失败 %v", err)
	}
	logrus.Infof("表结构迁移成功")
}
