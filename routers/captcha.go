package routers

import (
	"github.com/Xavier-Tse/lunar-gate/api"
	"github.com/gin-gonic/gin"
)

func CaptchaRouter(g *gin.RouterGroup) {
	r := g.Group("captcha")
	app := api.App.CaptchaApi
	r.GET("generate", app.Generate)
}
