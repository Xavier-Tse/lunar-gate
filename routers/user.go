package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/Xavier-Tse/lunar-gate/api"
	"github.com/Xavier-Tse/lunar-gate/middleware"
)

func UserRouter(g *gin.RouterGroup) {
	r := g.Group("user")
	app := api.App.UserApi
	r.POST("login", app.Login)
	r.POST("register", app.Register)
	r.PUT("password", middleware.Auth(), app.UpdatePassword)
	r.PUT("info", middleware.Auth(), app.UpdateUserinfo)
	r.GET("info", middleware.Auth(), app.Userinfo)
	r.GET("list", middleware.Auth(), app.List)
	r.DELETE("remove", middleware.Auth(), app.Remove)
}
