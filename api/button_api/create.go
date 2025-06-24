package button_api

import (
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	ID    uint   `json:"id"`
	Title string `json:"title" binding:"required,max=32"`
	Name  string `json:"name" binding:"required,max=32"`
	Group string `json:"group" binding:"required,max=32"`
}

func (ButtonApi) Create(c *gin.Context) {
	var cr CreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailBinding(err, c)
		return
	}
	var btn model.Button
	err := global.DB.Take(&btn, "name = ? or title = ?", cr.Name, cr.Title).Error
	if err == nil {
		res.FailWithMessage("name或title重复", c)
		return
	}

	btn = model.Button{
		Name:  cr.Name,
		Title: cr.Title,
		Group: cr.Group,
	}
	err = global.DB.Create(&btn).Error

	if err != nil {
		res.FailWithMessage("按钮创建失败", c)
		return
	}
	res.OkWithData(btn.ID, c)
}
