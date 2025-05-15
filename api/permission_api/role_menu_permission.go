package permission_api

import (
	"fmt"
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/Xavier-Tse/lunar-gate/utils/set"
	"github.com/gin-gonic/gin"
)

type RoleMenuPermissionRequest struct {
	RoleID     uint   `json:"roleID" binding:"required"`
	MenuIDList []uint `json:"menuIDList" binding:"required"`
}

func (PermissionApi) RoleMenuPermission(c *gin.Context) {
	var cr RoleMenuPermissionRequest
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

	var menuList []model.Menu
	global.DB.Find(&menuList, "id in ?", cr.MenuIDList)
	if len(menuList) != len(cr.MenuIDList) {
		res.FailWithMessage("选中的menu id不可用，请检查", c)
		return
	}

	for _, menu := range menuList {
		if !menu.Enable {
			res.FailWithMessage("选中隐藏菜单，请检查", c)
			return
		}
	}

	var menuRoleIDList []uint
	global.DB.Model(model.RoleMenu{}).
		Where("role_id = ?", cr.RoleID).
		Select("menu_id").Scan(&menuRoleIDList)

	intersectList := set.IntersectArray(menuRoleIDList, cr.MenuIDList)
	removeList := set.DiffArray(menuRoleIDList, intersectList)
	addList := set.DiffArray(cr.MenuIDList, intersectList)
	if len(addList) > 0 {
		var addRoleMenuList []model.RoleMenu
		for _, i2 := range addList {
			addRoleMenuList = append(addRoleMenuList, model.RoleMenu{
				MenuID: i2,
				RoleID: cr.RoleID,
			})
		}
		global.DB.Create(&addRoleMenuList)
	}

	if len(removeList) > 0 {
		var removeRoleMenuList []model.RoleMenu
		global.DB.Find(&removeRoleMenuList, "role_id = ? and menu_id in ?", cr.RoleID, removeList)
		global.DB.Delete(&removeRoleMenuList)
	}
	msg := fmt.Sprintf("新增 %d 个，删除 %d 个", len(addList), len(removeList))
	res.OkWithMessage(msg, c)
}
