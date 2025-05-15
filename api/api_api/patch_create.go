package api_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/model"
)

type apiInfo struct {
	Name   string `json:"name" binding:"required,max=32"`
	Path   string `json:"path" binding:"required,max=255"`
	Method string `json:"method" binding:"required,max=16"`
	Group  string `json:"group" binding:"required,max=32"`
}

type PatchCreateRequest struct {
	List []apiInfo `json:"list" binding:"required,dive"`
}

func (ApiApi) PatchCreate(c *gin.Context) {
	var cr PatchCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailBinding(err, c)
		return
	}

	var apiList []model.Api
	global.DB.Find(&apiList)
	var mpsMP = map[string]bool{} // method path
	var mpsN = map[string]bool{}  // name
	for _, api := range apiList {
		key := fmt.Sprintf("%s-%s", api.Method, api.Path)
		mpsMP[key] = true
		mpsN[api.Name] = true
	}

	var createApiList []model.Api
	for _, info := range cr.List {
		// 把重复的去掉
		key := fmt.Sprintf("%s-%s", info.Method, info.Path)
		if mpsMP[key] {
			continue
		}
		if mpsN[info.Name] {
			continue
		}
		createApiList = append(createApiList, model.Api{
			Name:   info.Name,
			Path:   info.Path,
			Method: info.Method,
			Group:  info.Group,
		})
	}
	if len(createApiList) > 0 {
		err := global.DB.Create(&createApiList).Error
		if err != nil {
			res.FailWithMessage("批量创建api失败", c)
			return
		}
	}
	msg := fmt.Sprintf("批量创建api记录成功，成功 %d 个", len(createApiList))
	res.OkWithMessage(msg, c)
}
