package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lunarise-dev/lunar-gate/common/res"
	"github.com/lunarise-dev/lunar-gate/global"
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

	var list = make([]model.User, 0)
	offset := (cr.Page.Page - 1) * cr.Limit
	global.DB.Preload("RoleList").Where(model.User{
		Username: cr.Username,
		Email:    cr.Email,
	}).Order(cr.Sort).Where("nickname like ?", fmt.Sprintf("%%%s%%", cr.Key)).Limit(cr.Limit).Offset(offset).Find(&list)

	var count int64
	global.DB.Model(model.User{}).Where(model.User{
		Username: cr.Username,
		Email:    cr.Email,
	}).Count(&count)
	res.OkWithData(gin.H{"list": list, "count": count}, c)
}
