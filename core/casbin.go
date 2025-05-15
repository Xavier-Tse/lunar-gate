package core

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/sirupsen/logrus"
)

func InitCasbin() *casbin.CachedEnforcer {
	// STEP 1: 创建 Gorm 适配器
	a, err := gormadapter.NewAdapterByDB(global.DB)
	if err != nil {
		logrus.Fatalf("创建Casbin GORM适配器失败: %v", err)
	}

	// STEP 2: 定义 Casbin 模型
	casbinModel := `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
`
	m, err := model.NewModelFromString(casbinModel)
	if err != nil {
		logrus.Fatalf("casbin模型加载失败: %v", err)
	}

	// STEP 3: 创建enforcer
	e, err := casbin.NewCachedEnforcer(m, a)
	if err != nil {
		logrus.Fatalf("创建Casbin Enforcer失败: %v", err)
	}

	// STEP 4: 设置缓存时间
	e.SetExpireTime(60 * 60)

	// STEP 5: 加载策略
	if err := e.LoadPolicy(); err != nil {
		logrus.Fatalf("Casbin策略加载失败: %v", err)
	}
	return e
}
