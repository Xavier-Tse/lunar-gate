package routers

import (
	"github.com/Xavier-Tse/lunar-gate/api"
	"github.com/gin-gonic/gin"
)

func WsRouter(g *gin.RouterGroup) {
	r := g.Group("ws")
	app := api.App.WsApi
	r.GET("", app.Ws)
}
