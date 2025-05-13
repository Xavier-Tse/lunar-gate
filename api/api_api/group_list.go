package api_api

import (
	"github.com/gin-gonic/gin"
	"github.com/lunarise-dev/lunar-gate/common/res"
	"github.com/lunarise-dev/lunar-gate/global"
	"github.com/lunarise-dev/lunar-gate/model"
)

func (ApiApi) GroupList(c *gin.Context) {
	var data = map[string][]model.Api{}
	var list []model.Api
	global.DB.Find(&list)
	for _, u := range list {
		data[u.Group] = append(data[u.Group], u)
	}
	res.OkWithData(data, c)
}
