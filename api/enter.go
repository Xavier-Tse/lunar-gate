package api

import (
	"github.com/Xavier-Tse/lunar-gate/api/api_api"
	"github.com/Xavier-Tse/lunar-gate/api/captcha_api"
	"github.com/Xavier-Tse/lunar-gate/api/email_api"
	"github.com/Xavier-Tse/lunar-gate/api/menu_api"
	"github.com/Xavier-Tse/lunar-gate/api/permission_api"
	"github.com/Xavier-Tse/lunar-gate/api/role_api"
	"github.com/Xavier-Tse/lunar-gate/api/user_api"
)

type Api struct {
	user_api.UserApi
	captcha_api.CaptchaApi
	email_api.EmailApi
	menu_api.MenuApi
	role_api.RoleApi
	api_api.ApiApi
	permission_api.PermissionApi
}

var App = new(Api)
