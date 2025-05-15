package user_api

import (
	"github.com/gin-gonic/gin"
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/middleware"
	"github.com/Xavier-Tse/lunar-gate/model"
)

type UpdateUserinfoRequest struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

func (UserApi) UpdateUserinfo(c *gin.Context) {
	var cr UpdateUserinfoRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailBinding(err, c)
		return
	}
	claims := middleware.GetAuth(c)

	var user model.User
	err := global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	err = global.DB.Model(&user).Updates(model.User{
		Nickname: cr.Nickname,
		Avatar:   cr.Avatar,
	}).Error
	if err != nil {
		res.FailWithMessage("用户信息修改失败", c)
		return
	}
	res.OkWithMessage("用户信息修改成功", c)
}
