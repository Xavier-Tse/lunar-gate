package menu_api

import (
	"github.com/gin-gonic/gin"
	"github.com/lunarise-dev/lunar-gate/common/res"
	"github.com/lunarise-dev/lunar-gate/global"
	"github.com/lunarise-dev/lunar-gate/model"
)

type MenuCreateRequest struct {
	Icon         string `json:"icon"`
	Title        string `json:"title" binding:"required"`
	Enable       bool   `json:"enable"`
	Name         string `json:"name" binding:"required"`
	Path         string `json:"path" binding:"required"`
	Component    string `json:"component"`
	ParentMenuID *uint  `json:"parentMenuID"`
	Sort         int    `json:"sort"`
}

func (MenuApi) MenuCreate(c *gin.Context) {
	var cr MenuCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailBinding(err, c)
		return
	}

	var menu model.Menu
	err := global.DB.Take(&menu, "name = ?", cr.Name).Error
	if err == nil {
		res.FailWithMessage("菜单名称不能重复", c)
		return
	}
	if cr.ParentMenuID != nil {
		var parent model.Menu
		err = global.DB.Take(&parent, *cr.ParentMenuID).Error
		if err != nil {
			res.FailWithMessage("父菜单不存在", c)
			return
		}
	}

	menu = model.Menu{
		Enable:    cr.Enable,
		Name:      cr.Name,
		Path:      cr.Path,
		Component: cr.Component,
		Meta: model.Meta{
			Icon:  cr.Icon,
			Title: cr.Title,
		},
		ParentMenuID: cr.ParentMenuID,
		Sort:         cr.Sort,
	}
	err = global.DB.Create(&menu).Error
	if err != nil {
		res.FailWithMessage("菜单创建失败", c)
		return
	}

	res.Ok(menu.ID, "菜单创建成功", c)
}
