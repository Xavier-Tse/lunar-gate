package menu_api

import (
	"github.com/gin-gonic/gin"
	"github.com/lunarise-dev/lunar-gate/common/query"
	"github.com/lunarise-dev/lunar-gate/common/res"
	"github.com/lunarise-dev/lunar-gate/global"
	"github.com/lunarise-dev/lunar-gate/model"
)

type MenuListRequest struct {
	model.Page
}

func (MenuApi) MenuList(c *gin.Context) {
	var cr MenuListRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailBinding(err, c)
		return
	}

	cr.Page.Sort = "sort desc"
	list, count, _ := query.List(model.Menu{}, query.Option{
		Page:     cr.Page,
		Likes:    []string{"name", "title"},
		Where:    global.DB.Where("parent_menu_id is null"),
		Preloads: []string{"Children"},
	})
	res.OkWithList(list, count, c)
}
