package role_api

import (
	"github.com/Xavier-Tse/lunar-gate/common/query"
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/gin-gonic/gin"
)

type RoleListRequest struct {
	model.Page
}

type RoleListResponse struct {
	model.LunarModel
	Title         string `json:"title"`
	RoleUserCount int    `json:"roleUserCount"`
	RoleApiCount  int    `json:"roleApiCount"`
	RoleMenuCount int    `json:"roleMenuCount"`
	ApiIDList     []uint `json:"apiIDList"`
}

func (RoleApi) List(c *gin.Context) {
	var cr RoleListRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailBinding(err, c)
		return
	}

	_list, count, _ := query.List(model.Role{}, query.Option{
		Page:     cr.Page,
		Likes:    []string{"title"},
		Preloads: []string{"UserList", "MenuList", "ApiList"},
	})
	var list = make([]RoleListResponse, 0)
	for _, role := range _list {
		var apiIDList = make([]uint, 0)
		for _, api := range role.ApiList {
			apiIDList = append(apiIDList, api.ID)
		}
		list = append(list, RoleListResponse{
			LunarModel:    role.LunarModel,
			Title:         role.Title,
			RoleUserCount: len(role.UserList),
			RoleMenuCount: len(role.MenuList),
			RoleApiCount:  len(role.ApiList),
			ApiIDList:     apiIDList,
		})
	}
	res.OkWithList(list, count, c)
}
