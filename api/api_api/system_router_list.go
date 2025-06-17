package api_api

import (
	"fmt"
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/gin-gonic/gin"
)

type RouterInfo struct {
	Method string `json:"method"`
	Path   string `json:"path"`
	Group  string `json:"group"`
	Name   string `json:"name"`
	Enable bool   `json:"enable"` // 是否同步
}

func (ApiApi) SystemRouterList(c *gin.Context) {
	var list = make([]RouterInfo, 0)

	// 数据库中有记录就标记为启用
	var apiList []model.Api
	global.DB.Find(&apiList)
	var mps = map[string]bool{}
	for _, api := range apiList {
		key := fmt.Sprintf("%s-%s", api.Method, api.Path)
		mps[key] = true
	}

	for _, i2 := range global.RoutesInfo {
		key := fmt.Sprintf("%s-%s", i2.Method, i2.Path)
		list = append(list, RouterInfo{
			Method: i2.Method,
			Path:   i2.Path,
			Enable: mps[key],
		})
	}
	res.OkWithData(list, c)
}
