package button_api

import (
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/gin-gonic/gin"
)

type UpdateRequest struct {
	ID    uint   `json:"id"`
	Title string `json:"title" binding:"required,max=32"`
	Name  string `json:"name" binding:"required,max=32"`
	Group string `json:"group" binding:"required,max=32"`
}

func (ButtonApi) Update(c *gin.Context) {
	var cr UpdateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailBinding(err, c)
		return
	}
	if cr.ID == 0 {
		res.FailWithMessage("请选择更新的按钮", c)
		return
	}

	var btn model.Button
	err := global.DB.Take(&btn, cr.ID).Error
	if err != nil {
		res.FailWithMessage("按钮不存在", c)
		return
	}

	var oldModel model.Button
	err = global.DB.Not("id = ?", btn.ID).Take(&oldModel, "name = ? or title = ?", cr.Name, cr.Title).Error
	if err == nil {
		res.FailWithMessage("name或title重复", c)
		return
	}

	err = global.DB.Model(&btn).Updates(model.Button{
		Name:  cr.Name,
		Title: cr.Title,
		Group: cr.Group,
	}).Error

	if err != nil {
		res.FailWithMessage("按钮更新失败", c)
		return
	}
	res.OkWithMessage("按钮更新成功", c)
}
