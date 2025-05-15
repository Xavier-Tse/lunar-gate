package routers

import (
	"github.com/Xavier-Tse/lunar-gate/api"
	"github.com/gin-gonic/gin"
)

func UserRouter(g *gin.RouterGroup) {
	r := g.Group("user")
	app := api.App.UserApi
	r.POST("login", app.Login)
	r.POST("register", app.Register)
	r.PUT("password", app.UpdatePassword)
	r.PUT("info", app.UpdateUserinfo)
	r.GET("info", app.Userinfo)
	r.GET("list", app.List)
	r.DELETE("remove", app.Remove)
	r.PUT("role", app.UserRoleUpdate)
}
