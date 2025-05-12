package menu_api

import (
	"github.com/gin-gonic/gin"
	"github.com/lunarise-dev/lunar-gate/common/res"
	"github.com/lunarise-dev/lunar-gate/global"
	"github.com/lunarise-dev/lunar-gate/model"
)

func (MenuApi) Tree(c *gin.Context) {
	var menuList []*model.Menu
	global.DB.Order("sort desc").Find(&menuList, "parent_menu_id is null")
	for _, u := range menuList {
		findSubMenuList(u)
	}
	res.OkWithData(menuList, c)
}

func findSubMenuList(model *model.Menu) {
	global.DB.Preload("Children").Take(&model)
	for _, child := range model.Children {
		findSubMenuList(child)
	}
}
