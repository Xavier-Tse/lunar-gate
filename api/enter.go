package api

import (
	"github.com/lunarise-dev/lunar-gate/api/captcha_api"
	"github.com/lunarise-dev/lunar-gate/api/email_api"
	"github.com/lunarise-dev/lunar-gate/api/user_api"
)

type Api struct {
	user_api.UserApi
	captcha_api.CaptchaApi
	email_api.EmailApi
}

var App = new(Api)
