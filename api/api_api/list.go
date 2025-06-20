package api_api

import (
	"github.com/Xavier-Tse/lunar-gate/common/query"
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/gin-gonic/gin"
)

type ApiListRequest struct {
	model.Page
	Method string `form:"method"`
	Group  string `form:"group"`
}

func (ApiApi) List(c *gin.Context) {
	var cr ApiListRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailBinding(err, c)
		return
	}

	list, count, _ := query.List(model.Api{
		Method: cr.Method,
		Group:  cr.Group,
	}, query.Option{
		Page:  cr.Page,
		Likes: []string{"name", "path"},
	})
	res.OkWithList(list, count, c)
}
