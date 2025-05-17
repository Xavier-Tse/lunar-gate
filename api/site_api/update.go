package site_api

import (
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/config"
	"github.com/Xavier-Tse/lunar-gate/core"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/gin-gonic/gin"
)

func (SiteApi) SiteUpdate(c *gin.Context) {
	name := c.Param("name")
	switch name {
	case "site":
		var cr config.Site
		err := c.ShouldBindJSON(&cr)
		if err != nil {
			res.FailBinding(err, c)
			return
		}
		global.Config.Site = cr
	case "email":
		var cr config.Email
		err := c.ShouldBindJSON(&cr)
		if err != nil {
			res.FailBinding(err, c)
			return
		}
		if cr.AuthCode == "******" {
			cr.AuthCode = global.Config.Email.AuthCode
		}
		global.Config.Email = cr
	case "jwt":
		var cr config.Jwt
		err := c.ShouldBindJSON(&cr)
		if err != nil {
			res.FailBinding(err, c)
			return
		}
		if cr.Secret == "******" {
			cr.Secret = global.Config.Jwt.Secret
		}
		global.Config.Jwt = cr
	default:
		res.FailWithMessage("站点配置名称错误", c)
		return
	}
	core.SetConfig(global.Config)
	res.OkWithMessage("修改配置成功", c)
}
