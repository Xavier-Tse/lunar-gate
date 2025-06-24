package button_api

import (
	"github.com/Xavier-Tse/lunar-gate/common/query"
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/gin-gonic/gin"
)

type ButtonListRequest struct {
	model.Page
	Group string `form:"group"`
}

func (ButtonApi) List(c *gin.Context) {
	var cr ButtonListRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailBinding(err, c)
		return
	}
	list, count, _ := query.List(model.Button{
		Group: cr.Group,
	}, query.Option{
		Page:  cr.Page,
		Likes: []string{"name", "title"},
	})
	res.OkWithList(list, count, c)
}
