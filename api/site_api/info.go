package site_api

import (
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/gin-gonic/gin"
)

func (SiteApi) Info(c *gin.Context) {
	// /api/site/site
	// /api/site/email
	// /api/site/jwt

	name := c.Param("name")
	var data any
	switch name {
	case "site":
		data = global.Config.Info.Site
	case "project":
		data = global.Config.Info.Project
	case "captcha":
		data = global.Config.Info.Login.Captcha
	case "email":
		d := global.Config.Info.Login.Email
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
