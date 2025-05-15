package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/model"
)

func (UserApi) Remove(c *gin.Context) {
	var cr model.IDListRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailBinding(err, c)
		return
	}

	var list []model.User
	global.DB.Find(&list, "id in ?", cr.IDList)
	var count int64
	if len(list) > 0 {
		count = global.DB.Delete(&list).RowsAffected

	}
	msg := fmt.Sprintf("删除用户 %d 个", count)
	res.OkWithMessage(msg, c)
}
