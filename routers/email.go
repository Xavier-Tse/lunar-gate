package routers

import (
	"github.com/Xavier-Tse/lunar-gate/api"
	"github.com/gin-gonic/gin"
)

func EmailRouter(g *gin.RouterGroup) {
	r := g.Group("email")
	app := api.App.EmailApi
	r.POST("send", app.Send)
}
