package user_api

import (
	"fmt"
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/Xavier-Tse/lunar-gate/utils/set"
	"github.com/gin-gonic/gin"
)

type UserRoleUpdateRequest struct {
	UserID     uint   `json:"userID" binding:"required"`
	RoleIDList []uint `json:"roleIDList" binding:"required"`
}

func (UserApi) UserRoleUpdate(c *gin.Context) {
	var cr UserRoleUpdateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailBinding(err, c)
		return
	}

	var user model.User
	err := global.DB.Take(&user, cr.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}

	var roleList []model.Role
	global.DB.Find(&roleList, "id in ?", cr.RoleIDList)
	if len(roleList) != len(cr.RoleIDList) {
		res.FailWithMessage("选中的role id不可用，请检查", c)
		return
	}

	var userRoleIDList []uint
	global.DB.Model(model.UserRole{}).
		Where("user_id = ?", cr.UserID).
		Select("role_id").Scan(&userRoleIDList)

	// 找出要添加的和要删除的
	// 已有 1, 2, 3, 4
	// 传入 3, 4, 5, 6
	// 添加 5, 6
	// 删除 1, 2
	intersectList := set.IntersectArray(userRoleIDList, cr.RoleIDList)
	removeList := set.DiffArray(userRoleIDList, intersectList)
	addList := set.DiffArray(cr.RoleIDList, intersectList)
	if len(addList) > 0 {
		var addUserRoleList []model.UserRole
		for _, i2 := range addList {
			addUserRoleList = append(addUserRoleList, model.UserRole{
				UserID: cr.UserID,
				RoleID: i2,
			})
		}
		global.DB.Create(&addUserRoleList)
	}

	if len(removeList) > 0 {
		var removeUserRoleList []model.UserRole
		global.DB.Find(&removeUserRoleList, "user_id = ? and role_id in ?", cr.UserID, removeList)
		global.DB.Unscoped().Delete(&removeUserRoleList)
	}
	msg := fmt.Sprintf("新增角色 %d 个，删除角色 %d 个", len(addList), len(removeList))
	res.OkWithMessage(msg, c)
}
