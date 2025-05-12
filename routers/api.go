package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lunarise-dev/lunar-gate/api"
	"github.com/lunarise-dev/lunar-gate/middleware"
)

func ApiRouter(g *gin.RouterGroup) {
	r := g.Group("api")
	app := api.App
	r.POST("create", middleware.Auth(), app.ApiCreate)
	r.PUT("update", middleware.Auth(), app.ApiUpdate)
	r.GET("list", middleware.Auth(), app.ApiList)
	r.DELETE("remove", middleware.Auth(), app.ApiRemove)
}
