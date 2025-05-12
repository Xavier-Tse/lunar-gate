package api_api

import (
	"github.com/gin-gonic/gin"
	"github.com/lunarise-dev/lunar-gate/common/res"
	"github.com/lunarise-dev/lunar-gate/global"
	"github.com/lunarise-dev/lunar-gate/model"
)

type ApiCreateRequest struct {
	Name   string `json:"name" binding:"required,max=32"`
	Path   string `json:"path" binding:"required,max=255"`
	Method string `json:"method" binding:"required,max=16"`
	Group  string `json:"group" binding:"required,max=32"` // api分组
}

func (ApiApi) ApiCreate(c *gin.Context) {
	var cr ApiCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailBinding(err, c)
		return
	}

	var api model.Api
	err := global.DB.Take(&api, "name = ?", cr.Name).Error
	if err == nil {
		res.FailWithMessage("api名称不能重复", c)
		return
	}

	api = model.Api{
		Name:   cr.Name,
		Path:   cr.Path,
		Method: cr.Method,
		Group:  cr.Group,
	}
	err = global.DB.Create(&api).Error

	if err != nil {
		res.FailWithMessage("api记录创建失败", c)
		return
	}
	res.OkWithData(api.ID, c)
}
