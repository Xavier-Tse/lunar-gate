package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lunarise-dev/lunar-gate/api"
)

func CaptchaRouter(g *gin.RouterGroup) {
	r := g.Group("captcha")
	app := api.App.CaptchaApi
	r.GET("/generate", app.Generate)
}
