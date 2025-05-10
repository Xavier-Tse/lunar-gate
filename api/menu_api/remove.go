package menu_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lunarise-dev/lunar-gate/common/res"
	"github.com/lunarise-dev/lunar-gate/global"
	"github.com/lunarise-dev/lunar-gate/model"
)

func (MenuApi) MenuRemove(c *gin.Context) {
	var cr model.IDListRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailBinding(err, c)
		return
	}

	var list []model.Menu
	global.DB.Find(&list, "id in ?", cr.IDList)
	var count int64
	if len(list) > 0 {
		count = global.DB.Delete(&list).RowsAffected

	}
	msg := fmt.Sprintf("删除菜单 %d 个", count)
	res.OkWithMessage(msg, c)
}
