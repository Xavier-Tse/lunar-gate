package routers

import (
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Run() {
	s := global.Config.System
	gin.SetMode(s.Mode)
	r := gin.Default()
	g := r.Group("api")

	g.Use(middleware.Auth())
	g.Use(middleware.Permission())

	UserRouter(g)
	CaptchaRouter(g)
	EmailRouter(g)
	MenuRouter(g)
	RoleRouter(g)
	ApiRouter(g)
	PermissionRouter(g)
	DataRouter(g)
	SiteRouter(g)
	WsRouter(g)

	r.Static("/static", "./static")

	global.RoutesInfo = r.Routes()

	logrus.Infof("web服务运行在 %s", s.Addr())
	err := r.Run(s.Addr())
	if err != nil {
		logrus.Fatalf("web服务启动失败 %s", err.Error())
	}
}
