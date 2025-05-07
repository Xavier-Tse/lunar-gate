package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lunarise-dev/lunar-gate/api"
)

func UserRouter(g *gin.RouterGroup) {
	r := g.Group("user")
	app := api.App
	r.POST("login", app.Login)
}
