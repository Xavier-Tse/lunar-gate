package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/Xavier-Tse/lunar-gate/api"
	"github.com/Xavier-Tse/lunar-gate/middleware"
)

func RoleRouter(g *gin.RouterGroup) {
	r := g.Group("role")
	app := api.App.RoleApi
	r.POST("create", middleware.Auth(), app.Create)
	r.GET("list", middleware.Auth(), app.List)
	r.PUT("update", middleware.Auth(), app.Update)
	r.DELETE("remove", middleware.Auth(), app.Remove)
}
