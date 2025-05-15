package api_api

import (
	"fmt"
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ApiUpdateRequest struct {
	ID     uint   `json:"id"`
	Name   string `json:"name" binding:"required,max=32"`
	Path   string `json:"path" binding:"required,max=255"`
	Method string `json:"method" binding:"required,max=16"`
	Group  string `json:"group" binding:"required,max=32"` // api分组
}

func (ApiApi) Update(c *gin.Context) {
	var cr ApiUpdateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailBinding(err, c)
		return
	}

	if cr.ID == 0 {
		res.FailWithMessage("请选择需要更新的api记录", c)
		return
	}
	var api model.Api
	err := global.DB.Take(&api, cr.ID).Error
	if err != nil {
		res.FailWithMessage("api记录不存在", c)
		return
	}
	var oldModel model.Api
	err = global.DB.Not("id = ?", api.ID).Take(&oldModel, "name = ?", cr.Name).Error
	if err == nil {
		res.FailWithMessage("api名称不能重复", c)
		return
	}

	// 判断有没有更新 path和method
	if api.Path != cr.Path || api.Method != cr.Method {
		logrus.Infof("api发生变化 %s %s => %s %s", api.Method, api.Path, cr.Method, cr.Path)

		var roleApiList []model.RoleApi
		global.DB.Find(&roleApiList, "api_id = ?", cr.ID)
		for _, r := range roleApiList {
			oldPolicy := []string{fmt.Sprintf("%d", r.RoleID), api.Path, api.Method}
			newPolicy := []string{fmt.Sprintf("%d", r.RoleID), cr.Path, cr.Method}
			global.Casbin.UpdatePolicy(oldPolicy, newPolicy)
		}
		if len(roleApiList) > 0 {
			global.Casbin.SavePolicy()
		}
	}

	err = global.DB.Model(&api).Updates(model.Api{
		Name:   cr.Name,
		Path:   cr.Path,
		Method: cr.Method,
		Group:  cr.Group,
	}).Error

	if err != nil {
		res.FailWithMessage("api记录更新失败", c)
		return
	}
	res.OkWithMessage("api记录更新成功", c)
}
