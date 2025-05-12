package api_api

import (
	"github.com/gin-gonic/gin"
	"github.com/lunarise-dev/lunar-gate/common/query"
	"github.com/lunarise-dev/lunar-gate/common/res"
	"github.com/lunarise-dev/lunar-gate/model"
)

type ApiListRequest struct {
	model.Page
	Method string `form:"method"`
	Group  string `form:"group"`
}

func (ApiApi) ApiList(c *gin.Context) {
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
