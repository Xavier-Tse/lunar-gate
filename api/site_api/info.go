package site_api

import (
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/gin-gonic/gin"
)

func (SiteApi) Info(c *gin.Context) {
	// /api/site/info
	// /api/site/email
	// /api/site/jwt

	name := c.Param("name")
	var data any
	switch name {
	case "info":
		data = global.Config.Info
	case "email":
		d := global.Config.Email
		d.AuthCode = "******"
		data = d
	case "jwt":
		d := global.Config.Jwt
		d.Secret = "******"
		data = d
	default:
		res.FailWithMessage("站点配置名称错误", c)
		return
	}
	res.OkWithData(data, c)
}
