package routers

import (
	"github.com/Xavier-Tse/lunar-gate/api"
	"github.com/gin-gonic/gin"
)

func MenuRouter(g *gin.RouterGroup) {
	r := g.Group("menu")
	app := api.App.MenuApi
	r.POST("create", app.Create)
	r.GET("list", app.List)
	r.GET("tree", app.Tree)
	r.PUT("update", app.Update)
	r.DELETE("remove", app.Remove)
}
