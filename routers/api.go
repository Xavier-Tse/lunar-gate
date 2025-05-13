package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lunarise-dev/lunar-gate/api"
	"github.com/lunarise-dev/lunar-gate/middleware"
)

func ApiRouter(g *gin.RouterGroup) {
	r := g.Group("api")
	app := api.App.ApiApi
	r.POST("create", middleware.Auth(), app.Create)
	r.PUT("update", middleware.Auth(), app.Update)
	r.GET("list", middleware.Auth(), app.List)
	r.DELETE("remove", middleware.Auth(), app.Remove)
	r.GET("system", middleware.Auth(), app.SystemRouterList)
	r.POST("create/patch", middleware.Auth(), app.PatchCreate)
	r.GET("group", middleware.Auth(), app.GroupList)
}
