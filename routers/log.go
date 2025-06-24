package routers

import (
	"github.com/Xavier-Tse/lunar-gate/api"
	"github.com/gin-gonic/gin"
)

func LogRouter(g *gin.RouterGroup) {
	r := g.Group("log")
	app := api.App.LogApi
	r.GET("", app.List)
}
