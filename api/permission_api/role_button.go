package permission_api

import (
	"fmt"
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/Xavier-Tse/lunar-gate/utils/set"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RoleButtonPermissionRequest struct {
	RoleID uint   `json:"roleID" binding:"required"`
	IDList []uint `json:"idList" binding:"required"`
}

func (PermissionApi) RoleButton(c *gin.Context) {
	var cr RoleButtonPermissionRequest
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

	var btnList []model.Button
	global.DB.Find(&btnList, "id in ?", cr.IDList)
	if len(btnList) != len(cr.IDList) {
		res.FailWithMessage("按钮数量不一致，请检查", c)
		return
	}

	var btnRoleIDList []uint
	global.DB.Model(model.RoleButton{}).
		Where("role_id = ?", cr.RoleID).
		Select("button_id").Scan(&btnRoleIDList)

	intersectList := set.IntersectArray(btnRoleIDList, cr.IDList)
	removeList := set.DiffArray(btnRoleIDList, intersectList)
	addList := set.DiffArray(cr.IDList, intersectList)
	if len(addList) > 0 {
		var addRoleBtnList []model.RoleButton
		for _, i2 := range addList {
			addRoleBtnList = append(addRoleBtnList, model.RoleButton{
				ButtonID: i2,
				RoleID:   cr.RoleID,
			})
		}
		global.DB.Create(&addRoleBtnList)
		logrus.Infof("创建角色按钮 %d 个", len(addList))
	}

	if len(removeList) > 0 {
		var removeRoleBtnList []model.RoleButton
		global.DB.Find(&removeRoleBtnList, "role_id = ? and button_id in ?", cr.RoleID, removeList)
		global.DB.Unscoped().Delete(&removeRoleBtnList)
		logrus.Infof("创建角色按钮 %d 个", len(removeRoleBtnList))
	}

	msg := fmt.Sprintf("按钮 新增 %d 个，删除 %d 个", len(addList), len(removeList))
	logrus.Infof(msg)
	res.OkWithMessage(msg, c)
}
