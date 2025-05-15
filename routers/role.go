package routers

import (
	"github.com/Xavier-Tse/lunar-gate/api"
	"github.com/gin-gonic/gin"
)

func RoleRouter(g *gin.RouterGroup) {
	r := g.Group("role")
	app := api.App.RoleApi
	r.POST("create", app.Create)
	r.GET("list", app.List)
	r.PUT("update", app.Update)
	r.DELETE("remove", app.Remove)
}
