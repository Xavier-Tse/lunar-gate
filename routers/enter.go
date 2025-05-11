package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lunarise-dev/lunar-gate/global"
	"github.com/sirupsen/logrus"
)

func Run() {
	s := global.Config.System
	gin.SetMode(s.Mode)
	r := gin.Default()
	g := r.Group("api")

	UserRouter(g)
	CaptchaRouter(g)
	EmailRouter(g)
	MenuRouter(g)
	RoleRouter(g)

	r.Static("/static", "./static")

	logrus.Infof("web服务运行在 %s", s.Addr())
	err := r.Run(s.Addr())
	if err != nil {
		logrus.Fatalf("web服务启动失败 %s", err.Error())
	}
}
