package routers

import (
	"github.com/Xavier-Tse/lunar-gate/api"
	"github.com/gin-gonic/gin"
)

func UserRouter(g *gin.RouterGroup) {
	r := g.Group("user")
	app := api.App.UserApi
	r.POST("login", app.Login)
	r.POST("logout", app.Logout)
	r.POST("register", app.Register)
	r.PUT("password", app.UpdatePassword)
	r.PUT("info", app.UpdateUserinfo)
	r.GET("info", app.Userinfo)
	r.GET("", app.List)
	r.DELETE("", app.Remove)
	r.PUT("role", app.UserRoleUpdate)
}
