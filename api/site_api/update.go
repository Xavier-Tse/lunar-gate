package site_api

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/config"
	"github.com/Xavier-Tse/lunar-gate/core"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/gin-gonic/gin"
	"os"
)

func (SiteApi) Update(c *gin.Context) {
	name := c.Param("name")
	switch name {
	case "info":
		var cr config.Site
		err := c.ShouldBindJSON(&cr)
		if err != nil {
			res.FailBinding(err, c)
			return
		}
		global.Config.Info.Site = cr
	case "project":
		var cr config.Project
		err := c.ShouldBindJSON(&cr)
		if err != nil {
			res.FailBinding(err, c)
			return
		}
		if err = SetHtml(cr); err != nil {
			res.FailWithError(err, c)
			return
		}
		global.Config.Info.Project = cr
	case "captcha":
		var cr config.Captcha
		err := c.ShouldBindJSON(&cr)
		if err != nil {
			res.FailBinding(err, c)
			return
		}
		global.Config.Info.Login.Captcha = cr
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

func SetHtml(p config.Project) error {
	if p.Path == "" {
		return nil
	}
	f, err := os.Open(p.Path)
	if err != nil {
		return errors.New("前端文件不存在")
	}

	doc, _ := goquery.NewDocumentFromReader(f)
	if p.Title != "" {
		if doc.Find("title").Length() != 0 {
			doc.Find("title").SetText(p.Title)
		} else {
			doc.Find("head").AppendHtml(fmt.Sprintf("<title>%s</title>", p.Title))
		}
	}

	if p.Icon != "" {
		if doc.Find("link[rel=\"icon\"]").Length() != 0 {
			doc.Find("link[rel=\"icon\"]").SetAttr("href", p.Icon)
		} else {
			doc.Find("head").AppendHtml(fmt.Sprintf("<link rel=\"icon\" href=\"%s\">", p.Icon))
		}
	}

	ret, _ := doc.Html()
	err = os.WriteFile(p.Path, []byte(ret), 0666)
	return err
}
