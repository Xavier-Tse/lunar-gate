package permission_api

import (
	"fmt"
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/Xavier-Tse/lunar-gate/utils/set"
	"github.com/gin-gonic/gin"
)

type RoleApiPermissionRequest struct {
	RoleID    uint   `json:"roleID" binding:"required"`
	ApiIDList []uint `json:"apiIDList" binding:"required"`
}

func (PermissionApi) RoleApiPermission(c *gin.Context) {
	var cr RoleApiPermissionRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailBinding(err, c)
		return
	}

	var role model.Role
	err := global.DB.Take(&role, cr.RoleID).Error
	if err != nil {
		res.FailWithMessage("角色不存在", c)
		return
	}

	var apiList []model.Api
	global.DB.Find(&apiList, "id in ?", cr.ApiIDList)
	if len(apiList) != len(cr.ApiIDList) {
		res.FailWithMessage("选中的api id不可用，请检查", c)
		return
	}

	var apiRoleIDList []uint
	global.DB.Model(model.RoleApi{}).
		Where("role_id = ?", cr.RoleID).
		Select("api_id").Scan(&apiRoleIDList)

	// 找出要添加的和要删除的
	// 已有 1, 2, 3, 4
	// 传入 3, 4, 5, 6
	// 添加 5, 6
	// 删除 1, 2
	intersectList := set.IntersectArray(apiRoleIDList, cr.ApiIDList)
	removeList := set.DiffArray(apiRoleIDList, intersectList)
	addList := set.DiffArray(cr.ApiIDList, intersectList)
	if len(addList) > 0 {
		var addRoleApiList []model.RoleApi
		for _, i2 := range addList {
			addRoleApiList = append(addRoleApiList, model.RoleApi{
				ApiID:  i2,
				RoleID: cr.RoleID,
			})
		}
		global.DB.Create(&addRoleApiList)
		addRoleApiList = []model.RoleApi{}
		global.DB.Preload("Api").Find(&addRoleApiList, "role_id = ? and api_id in ?", cr.RoleID, addList)
		for _, api := range addRoleApiList {
			global.Casbin.AddPolicy(fmt.Sprintf("%d", api.RoleID), api.Api.Path, api.Api.Method)
		}
	}

	if len(removeList) > 0 {
		var removeRoleApiList []model.RoleApi
		global.DB.Preload("Api").Find(&removeRoleApiList, "role_id = ? and api_id in ?", cr.RoleID, removeList)
		for _, api := range removeRoleApiList {
			global.Casbin.RemovePolicy(fmt.Sprintf("%d", api.RoleID), api.Api.Path, api.Api.Method)
		}
		global.DB.Delete(&removeRoleApiList)
	}

	if len(removeList) > 0 {
		global.DB.Delete(&model.RoleApi{}, "role_id = ? and api_id in ?", cr.RoleID, removeList)
	}
	msg := fmt.Sprintf("新增 %d 个，删除 %d 个", len(addList), len(removeList))
	res.OkWithMessage(msg, c)
}
