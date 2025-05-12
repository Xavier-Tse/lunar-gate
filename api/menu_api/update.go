package menu_api

import (
	"github.com/gin-gonic/gin"
	"github.com/lunarise-dev/lunar-gate/common/res"
	"github.com/lunarise-dev/lunar-gate/global"
	"github.com/lunarise-dev/lunar-gate/model"
	"github.com/lunarise-dev/lunar-gate/utils/maps"
	"github.com/sirupsen/logrus"
)

type MenuUpdateRequest struct {
	ID           uint    `json:"id" binding:"required"`
	Icon         *string `json:"icon" maps:"icon"`
	Title        *string `json:"title" maps:"title"`
	Enable       *bool   `json:"enable" maps:"enable"`
	Name         *string `json:"name" maps:"name"`
	Path         *string `json:"path" maps:"path"`
	Component    *string `json:"component" maps:"component"`
	ParentMenuID *uint   `json:"parentMenuID" maps:"parent_menu_id"`
	Sort         *int    `json:"sort" maps:"sort"`
}

func (MenuApi) Update(c *gin.Context) {
	var cr MenuUpdateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailBinding(err, c)
		return
	}

	var menu model.Menu
	err := global.DB.Take(&menu, cr.ID).Error
	if err != nil {
		res.FailWithMessage("菜单不存在", c)
		return
	}

	mps := maps.Struct2Map(cr, "maps")
	if len(mps) == 0 {
		res.FailWithMessage("菜单无更新", c)
		return
	}

	name, ok := mps["name"]
	if ok {
		if name == "" {
			res.FailWithMessage("菜单name不能为空", c)
			return
		}

		var m model.Menu
		err = global.DB.Not("id = ?", menu.ID).Take(&m, "name = ?", name).Error
		if err == nil {
			res.FailWithMessage("菜单name不能重复", c)
			return
		}
	}

	path, ok := mps["path"]
	if ok {
		if path == "" {
			res.FailWithMessage("菜单path不能为空", c)
			return
		}
	}

	if cr.ParentMenuID == nil {
		mps["parent_menu_id"] = nil
	} else {
		parentMenuID := *cr.ParentMenuID
		mps["parent_menu_id"] = parentMenuID
		var m model.Menu
		err = global.DB.Take(&m, parentMenuID).Error
		if err != nil {
			res.FailWithMessage("父菜单不存在", c)
			return
		}

		subMenuList := model.FindSubMenuList(m)
		for _, menuModel := range subMenuList {
			if menuModel.ID == parentMenuID {
				res.FailWithMessage("菜单层级错误", c)
				return
			}
		}
	}

	logrus.Infof("菜单更新值：%v", mps)
	err = global.DB.Model(&menu).Updates(mps).Error
	if err != nil {
		res.FailWithMessage("菜单更新失败", c)
		return
	}
	res.Ok(menu.ID, "菜单更新成功", c)
}
