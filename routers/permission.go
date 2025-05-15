package routers

import (
	"github.com/Xavier-Tse/lunar-gate/api"
	"github.com/gin-gonic/gin"
)

func PermissionRouter(g *gin.RouterGroup) {
	r := g.Group("permission")
	app := api.App.PermissionApi
	r.PUT("api", app.RoleApiPermission)
}
