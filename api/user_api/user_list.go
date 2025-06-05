package user_api

import (
	"github.com/Xavier-Tse/lunar-gate/common/query"
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/gin-gonic/gin"
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

func (UserApi) List(c *gin.Context) {
	var cr UserListRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailBinding(err, c)
		return
	}

	list, count, _ := query.List(model.User{
		Username: cr.Username,
		Email:    cr.Email,
	}, query.Option{
		Page:     cr.Page,
		Likes:    []string{"nickname", "username"},
		Preloads: []string{"RoleList"},
	})
	res.OkWithList(list, count, c)
}
