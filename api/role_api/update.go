package role_api

import (
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/gin-gonic/gin"
)

type RoleUpdateRequest struct {
	ID    uint   `json:"id" binding:"required"`
	Title string `json:"title" binding:"required,max=16"`
}

func (RoleApi) Update(c *gin.Context) {
	var cr RoleUpdateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailBinding(err, c)
		return
	}

	var role model.Role
	err := global.DB.Take(&role, cr.ID).Error
	if err != nil {
		res.FailWithMessage("角色不存在", c)
		return
	}

	if cr.Title == role.Title {
		res.FailWithMessage("角色名称未修改", c)
		return
	}

	var m model.Role
	err = global.DB.Not("id = ?", cr.ID).Take(&m, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("角色名称重复", c)
		return
	}
	global.DB.Model(&role).Update("title", cr.Title)
	res.Ok(role.ID, "角色更新成功", c)
}
