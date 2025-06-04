package routers

import (
	"fmt"
	"github.com/Xavier-Tse/lunar-gate/api/ws_api"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

// defaultLogFormatter is the default log format function Logger middleware uses.
var defaultLogFormatter = func(param gin.LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	if param.Latency > time.Minute {
		param.Latency = param.Latency.Truncate(time.Second)
	}
	msg := fmt.Sprintf("[GIN] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, param.StatusCode, resetColor,
		param.Latency,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
		param.ErrorMessage,
	)
	ws_api.SendLogMsg(msg)
	return msg
}

func Run() {
	s := global.Config.System
	gin.SetMode(s.Mode)
	// 修改为gin.New()以避免默认的Logger中间件
	r := gin.New()

	// 使用自定义日志格式器并添加Recovery中间件
	r.Use(gin.Recovery(), gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: defaultLogFormatter,
	}))

	g := r.Group("api")

	g.Use(middleware.Auth())
	g.Use(middleware.Permission())
	//g.Use(middleware.Cors())

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
