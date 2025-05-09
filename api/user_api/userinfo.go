package user_api

import (
	"github.com/gin-gonic/gin"
	"github.com/lunarise-dev/lunar-gate/common/res"
	"github.com/lunarise-dev/lunar-gate/global"
	"github.com/lunarise-dev/lunar-gate/middleware"
	"github.com/lunarise-dev/lunar-gate/model"
)

type UserinfoResponse struct {
	UserID   uint   `json:"userID"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	RoleList []uint `json:"roleList"`
}

func (UserApi) Userinfo(c *gin.Context) {
	claims := middleware.GetAuth(c)

	var user model.User
	err := global.DB.Preload("RoleList").Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	data := UserinfoResponse{
		UserID:   user.ID,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		RoleList: user.GetRoleList(),
	}
	res.OkWithData(data, c)
}
