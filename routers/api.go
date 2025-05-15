package routers

import (
	"github.com/Xavier-Tse/lunar-gate/api"
	"github.com/gin-gonic/gin"
)

func ApiRouter(g *gin.RouterGroup) {
	r := g.Group("api")
	app := api.App.ApiApi
	r.POST("create", app.Create)
	r.PUT("update", app.Update)
	r.GET("list", app.List)
	r.DELETE("remove", app.Remove)
	r.GET("system", app.SystemRouterList)
	r.POST("create/patch", app.PatchCreate)
	r.GET("group", app.GroupList)
}
