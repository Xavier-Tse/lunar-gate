package flags

import (
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
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
		&model.UserLogin{},
		&model.UserRole{},
		&gormadapter.CasbinRule{},
	); err != nil {
		logrus.Fatalf("表结构迁移失败 %v", err)
	}
	logrus.Infof("表结构迁移成功")
}
