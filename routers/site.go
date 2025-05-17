package routers

import (
	"github.com/Xavier-Tse/lunar-gate/api"
	"github.com/gin-gonic/gin"
)

func SiteRouter(g *gin.RouterGroup) {
	r := g.Group("site")
	app := api.App.SiteApi
	r.GET(":name", app.Info)
}
