package button_api

import (
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/gin-gonic/gin"
)

func (ButtonApi) GroupOptions(c *gin.Context) {
	var list = make([]model.OptionsResponse[string], 0)
	global.DB.Model(model.Button{}).Where("`group` <> ''").Group("group").Select("`group` as value", "`group` as label").Scan(&list)
	res.OkWithData(list, c)
}
