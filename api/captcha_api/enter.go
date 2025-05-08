package captcha_api

import (
	"github.com/gin-gonic/gin"
	"github.com/lunarise-dev/lunar-gate/common/res"
	"github.com/lunarise-dev/lunar-gate/global"
	"github.com/lunarise-dev/lunar-gate/utils/captcha"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

type CaptchaApi struct {
}

type GenerateCaptchaResponse struct {
	CaptchaID string `json:"captchaID"`
	Captcha   string `json:"captcha"`
}

func (CaptchaApi) GenerateCaptcha(c *gin.Context) {
	var driver base64Captcha.Driver
	switch global.Config.Captcha.Type {
	case "string":
		driver = stringCaptcha()
	}

	cp := base64Captcha.NewCaptcha(driver, captcha.Store)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		logrus.Errorf("图片验证码生成失败 %s", err)
		res.FailWithMessage("验证码生成失败", c)
		return
	}

	res.OkWithData(GenerateCaptchaResponse{
		CaptchaID: id,
		Captcha:   b64s,
	}, c)
}

func stringCaptcha() *base64Captcha.DriverString {
	d := &base64Captcha.DriverString{
		Width:           200,
		Height:          60,
		NoiseCount:      2,
		ShowLineOptions: 4,
		Length:          6,
		Source:          "0123456789",
	}
	return d
}
