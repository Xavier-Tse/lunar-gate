package routers

import (
	"github.com/Xavier-Tse/lunar-gate/api"
	"github.com/gin-gonic/gin"
)

func DataRouter(g *gin.RouterGroup) {
	r := g.Group("data")
	app := api.App.DataApi
	r.GET("sum", app.Summary)
	r.GET("user", app.UserTimeAggregation)
	r.GET("computer", app.Computer)
}
