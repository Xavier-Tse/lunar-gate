package routers

import (
	"github.com/Xavier-Tse/lunar-gate/api"
	"github.com/gin-gonic/gin"
)

func ImageRouter(g *gin.RouterGroup) {
	r := g.Group("image")
	app := api.App.ImageApi
	r.POST("upload", app.Upload)
}
