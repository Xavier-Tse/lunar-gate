package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lunarise-dev/lunar-gate/api"
	"github.com/lunarise-dev/lunar-gate/middleware"
)

func RoleRouter(g *gin.RouterGroup) {
	r := g.Group("/role")
	app := api.App
	r.POST("create", middleware.Auth(), app.RoleCreate)
	r.GET("list", middleware.Auth(), app.RoleList)
	r.PUT("update", middleware.Auth(), app.RoleUpdate)
}
