package routers

import (
	"github.com/Xavier-Tse/lunar-gate/api"
	"github.com/gin-gonic/gin"
)

func MenuRouter(g *gin.RouterGroup) {
	r := g.Group("menu")
	app := api.App.MenuApi
	r.POST("", app.Create)
	r.GET("", app.List)
	r.GET("tree", app.Tree)
	r.PUT("", app.Update)
	r.DELETE("", app.Remove)
}
