package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/Xavier-Tse/lunar-gate/api"
	"github.com/Xavier-Tse/lunar-gate/middleware"
)

func PermissionRouter(g *gin.RouterGroup) {
	r := g.Group("permission")
	app := api.App.PermissionApi
	r.PUT("api", middleware.Auth(), app.RoleApiPermission)
}
