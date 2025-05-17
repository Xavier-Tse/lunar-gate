package data_api

import (
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/gin-gonic/gin"
)

type SummaryResponse struct {
	UserCount        int64 `json:"userCount"`
	NowLoginCount    int64 `json:"nowLoginCount"`
	NowRegisterCount int64 `json:"nowRegisterCount"`
}

func (DataApi) Summary(c *gin.Context) {
	var data SummaryResponse
	global.DB.Model(model.User{}).Count(&data.UserCount)
	global.DB.Model(model.User{}).Where("date(created_at) = date(now())").Count(&data.NowRegisterCount)
	global.DB.Model(model.UserLogin{}).Where("date(created_at) = date(now())").Count(&data.NowLoginCount)
	res.OkWithData(data, c)
}
