package api_api

import (
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/gin-gonic/gin"
)

func (ApiApi) Options(c *gin.Context) {
	var list = make([]model.OptionsResponse[string], 0)
	global.DB.Model(model.Api{}).Where("`group` <> ''").Group("group").Select("`group` as value", "`group` as label").Scan(&list)
	res.OkWithData(list, c)
}
