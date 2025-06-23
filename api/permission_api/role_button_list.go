package permission_api

import (
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/middleware"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/gin-gonic/gin"
)

func (PermissionApi) RoleButtonList(c *gin.Context) {
	claims := middleware.GetAuth(c)
	var nameList = make([]string, 0)
	if claims.IsAdmin {
		global.DB.Model(model.Button{}).Select("name").Scan(&nameList)
		res.OkWithData(nameList, c)
		return
	}

	var btnIDList []uint
	global.DB.Model(model.RoleButton{}).Where("role_id in ?", claims.RoleList).
		Select("btn_id").Scan(&btnIDList)
	global.DB.Model(model.Button{}).Where("id in ?", btnIDList).Select("name").Scan(&nameList)
	res.OkWithData(nameList, c)
}
