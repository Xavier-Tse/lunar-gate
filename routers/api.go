package routers

import (
	"github.com/Xavier-Tse/lunar-gate/api"
	"github.com/gin-gonic/gin"
)

func ApiRouter(g *gin.RouterGroup) {
	r := g.Group("api")
	app := api.App.ApiApi
	r.POST("", app.Create)
	r.PUT("", app.Update)
	r.GET("", app.List)
	r.DELETE("", app.Remove)
	r.GET("system", app.SystemRouterList)
	r.POST("create/batch", app.BatchCreate)
	r.GET("group", app.GroupList)
	r.GET("options", app.Options)
}
