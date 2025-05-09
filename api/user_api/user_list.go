package user_api

import (
	"github.com/gin-gonic/gin"
	"github.com/lunarise-dev/lunar-gate/common/query"
	"github.com/lunarise-dev/lunar-gate/common/res"
	"github.com/lunarise-dev/lunar-gate/model"
)

type UserListRequest struct {
	model.Page
	Role     uint   `form:"role"`
	Username string `form:"username"`
	Email    string `form:"email"`
}

type UserListResponse struct {
	model.User
}

func (UserApi) UserList(c *gin.Context) {
	var cr UserListRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailBinding(err, c)
		return
	}

	list, count, _ := query.List(model.User{
		Username: cr.Username,
		Email:    cr.Email,
	}, query.Option{
		Page:  cr.Page,
		Likes: []string{"nickname", "username"},
	})
	res.OkWithList(list, count, c)
}
