package role_api

import (
	"github.com/gin-gonic/gin"
	"github.com/lunarise-dev/lunar-gate/common/res"
	"github.com/lunarise-dev/lunar-gate/global"
	"github.com/lunarise-dev/lunar-gate/model"
)

type RoleCreateRequest struct {
	Title string `json:"title" binding:"required,max=16"`
}

func (RoleApi) Create(c *gin.Context) {
	var cr RoleCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailBinding(err, c)
		return
	}

	var role model.Role
	err := global.DB.Take(&role, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("角色名称重复", c)
		return
	}
	role.Title = cr.Title
	err = global.DB.Create(&role).Error
	if err != nil {
		res.FailWithMessage("角色创建失败", c)
		return
	}
	res.Ok(role.ID, "角色创建成功", c)
}
