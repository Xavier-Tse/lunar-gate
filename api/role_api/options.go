package role_api

import (
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/gin-gonic/gin"
)

func (RoleApi) Options(c *gin.Context) {
	var list = make([]model.OptionsResponse[uint], 0)
	global.DB.Model(model.Role{}).Order("created_at desc").Select("id as value", "title as label").Scan(&list)
	res.OkWithData(list, c)
}
