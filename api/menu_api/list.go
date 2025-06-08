package menu_api

import (
	"github.com/Xavier-Tse/lunar-gate/common/query"
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/gin-gonic/gin"
)

type MenuListRequest struct {
	model.Page
}

func (MenuApi) List(c *gin.Context) {
	var cr MenuListRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailBinding(err, c)
		return
	}

	cr.Page.Sort = "sort desc"
	list, count, _ := query.List(&model.Menu{}, query.Option{
		Page:  cr.Page,
		Likes: []string{"name", "title"},
		Where: global.DB.Where("parent_menu_id is null"),
		Callback: func(_list any) {
			list := _list.([]*model.Menu)
			for _, menu := range list {
				model.SubMenuList(menu)
			}
		},
	})
	res.OkWithList(list, count, c)
}
