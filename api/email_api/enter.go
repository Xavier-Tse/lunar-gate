package email_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lunarise-dev/lunar-gate/common/res"
	"github.com/lunarise-dev/lunar-gate/global"
	"github.com/lunarise-dev/lunar-gate/utils/captcha"
	"github.com/lunarise-dev/lunar-gate/utils/email"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

type EmailApi struct {
}

type SendEmailRequest struct {
	Email       string `json:"email" binding:"required,email"`
	CaptchaID   string `json:"captchaID"`
	CaptchaCode string `json:"captchaCode"`
}
type SendEmailResponse struct {
	EmailID string `json:"emailID"`
}

func (EmailApi) SendEmail(c *gin.Context) {
	var cr SendEmailRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailBinding(err, c)
		return
	}

	if !global.Config.Email.Enable {
		res.FailWithMessage("管理员未配置邮箱，暂无法注册", c)
		return
	}

	if global.Config.Captcha.Enable {
		if cr.CaptchaID == "" || cr.CaptchaCode == "" {
			res.FailWithMessage("请输入图片验证码", c)
			return
		}
		if !captcha.Store.Verify(cr.CaptchaID, cr.CaptchaCode, true) {
			res.FailWithMessage("图片验证码验证失败", c)
			return
		}
	}

	// 生成的是邮箱验证码
	var driver = base64Captcha.DriverString{
		Width:           200,
		Height:          60,
		NoiseCount:      2,
		ShowLineOptions: 4,
		Length:          4,
		Source:          "0123456789",
	}
	cp := base64Captcha.NewCaptcha(&driver, captcha.Store)
	id, _, code, err := cp.Generate()
	if err != nil {
		logrus.Errorf("邮箱验证码生成失败 %s", err)
		res.FailWithMessage("验证码生成失败", c)
		return
	}

	content := fmt.Sprintf("您正在进行用户注册，这是你的验证码 %s，5分钟内有效", code)
	err = email.SendEmail(cr.Email, "用户注册", content)
	if err != nil {
		logrus.Errorf("邮箱发送失败 %s", err)
		res.FailWithMessage("邮箱发送失败", c)
		return
	}

	res.OkWithData(SendEmailResponse{
		EmailID: id,
	}, c)
}
