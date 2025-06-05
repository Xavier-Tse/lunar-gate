package routers

import (
	"github.com/Xavier-Tse/lunar-gate/api"
	"github.com/gin-gonic/gin"
)

func RoleRouter(g *gin.RouterGroup) {
	r := g.Group("role")
	app := api.App.RoleApi
	r.POST("", app.Create)
	r.GET("", app.List)
	r.PUT("", app.Update)
	r.DELETE("", app.Remove)
}
