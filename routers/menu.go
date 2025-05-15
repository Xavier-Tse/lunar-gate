package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/Xavier-Tse/lunar-gate/api"
	"github.com/Xavier-Tse/lunar-gate/middleware"
)

func MenuRouter(g *gin.RouterGroup) {
	r := g.Group("menu")
	app := api.App.MenuApi
	r.POST("create", middleware.Auth(), app.Create)
	r.GET("list", middleware.Auth(), app.List)
	r.GET("tree", middleware.Auth(), app.Tree)
	r.PUT("update", middleware.Auth(), app.Update)
	r.DELETE("remove", middleware.Auth(), app.Remove)
}
