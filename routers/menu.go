package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lunarise-dev/lunar-gate/api"
	"github.com/lunarise-dev/lunar-gate/middleware"
)

func MenuRouter(g *gin.RouterGroup) {
	r := g.Group("menu")
	app := api.App
	r.POST("create", middleware.Auth(), app.MenuCreate)
	r.GET("list", middleware.Auth(), app.MenuList)
	r.GET("tree", middleware.Auth(), app.MenuTree)
	r.PUT("update", middleware.Auth(), app.MenuUpdate)
}
