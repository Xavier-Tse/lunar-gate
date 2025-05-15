package role_api

import (
	"github.com/gin-gonic/gin"
	"github.com/Xavier-Tse/lunar-gate/common/query"
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/model"
)

type RoleListRequest struct {
	model.Page
}

type RoleListResponse struct {
	model.LunarModel
	Title         string `json:"title"`
	RoleUserCount int    `json:"roleUserCount"`
	RoleApiCount  int    `json:"roleApiCount"` // TODO: casbin融合之后再做
	RoleMenuCount int    `json:"roleMenuCount"`
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
		Preloads: []string{"UserList", "MenuList"},
	})
	var list = make([]RoleListResponse, 0)
	for _, role := range _list {
		list = append(list, RoleListResponse{
			LunarModel:    role.LunarModel,
			Title:         role.Title,
			RoleUserCount: len(role.UserList),
			RoleMenuCount: len(role.MenuList),
		})
	}
	res.OkWithList(list, count, c)
}
