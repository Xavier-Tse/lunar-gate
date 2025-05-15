package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/Xavier-Tse/lunar-gate/api"
)

func EmailRouter(g *gin.RouterGroup) {
	r := g.Group("email")
	app := api.App.EmailApi
	r.POST("send", app.Send)
}
