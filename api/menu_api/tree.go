package menu_api

import (
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/gin-gonic/gin"
)

func (MenuApi) Tree(c *gin.Context) {
	var menuList []*model.Menu
	global.DB.Order("sort desc").Find(&menuList, "parent_menu_id is null")
	for _, u := range menuList {
		model.SubMenuList(u)
	}
	res.OkWithData(menuList, c)
}
