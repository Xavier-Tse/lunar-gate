package permission_api

import (
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/middleware"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/gin-gonic/gin"
)

func (PermissionApi) RoleMenuTree(c *gin.Context) {
	claims := middleware.GetAuth(c)

	var menuIDList []uint
	global.DB.Model(model.RoleMenu{}).
		Where("role_id in ?", claims.RoleList).Select("menu_id").Scan(&menuIDList)
	var menuIDMap = map[uint]bool{}
	for _, u := range menuIDList {
		menuIDMap[u] = true
	}

	var menuList []*model.Menu
	global.DB.Order("sort desc").
		Find(&menuList, "parent_menu_id is null and id in ?", menuIDList)
	for _, menu := range menuList {
		findSubMenuList(menu, menuIDMap)
	}
	res.OkWithData(menuList, c)
}

func findSubMenuList(menu *model.Menu, menuIDMap map[uint]bool) {
	global.DB.Preload("Children").Take(&menu)

	var list = make([]*model.Menu, 0)
	for _, child := range menu.Children {
		if !menuIDMap[child.ID] {
			continue
		}
		list = append(list, child)
		findSubMenuList(child, menuIDMap)
	}
	menu.Children = list
}
