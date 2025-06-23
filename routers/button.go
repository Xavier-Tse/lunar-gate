package routers

import (
	"github.com/Xavier-Tse/lunar-gate/api"
	"github.com/gin-gonic/gin"
)

func ButtonRouter(g *gin.RouterGroup) {
	r := g.Group("button")
	app := api.App.ButtonApi
	r.POST("", app.Create)
	r.PUT("", app.Update)
	r.GET("", app.List)
	r.GET("options", app.GroupOptions)
	r.DELETE("", app.Remove)
}
