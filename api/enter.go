package api

import "github.com/lunarise-dev/lunar-gate/api/user_api"

type Api struct {
	user_api.UserApi
}

var App = new(Api)
