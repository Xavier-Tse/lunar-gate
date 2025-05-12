package api

import (
	"github.com/lunarise-dev/lunar-gate/api/api_api"
	"github.com/lunarise-dev/lunar-gate/api/captcha_api"
	"github.com/lunarise-dev/lunar-gate/api/email_api"
	"github.com/lunarise-dev/lunar-gate/api/menu_api"
	"github.com/lunarise-dev/lunar-gate/api/role_api"
	"github.com/lunarise-dev/lunar-gate/api/user_api"
)

type Api struct {
	user_api.UserApi
	captcha_api.CaptchaApi
	email_api.EmailApi
	menu_api.MenuApi
	role_api.RoleApi
	api_api.ApiApi
}

var App = new(Api)
