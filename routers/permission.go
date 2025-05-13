package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lunarise-dev/lunar-gate/api"
	"github.com/lunarise-dev/lunar-gate/middleware"
)

func PermissionRouter(g *gin.RouterGroup) {
	r := g.Group("permission")
	app := api.App.PermissionApi
	r.PUT("api", middleware.Auth(), app.RoleApiPermission)
}
