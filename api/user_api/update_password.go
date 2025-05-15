package user_api

import (
	"github.com/gin-gonic/gin"
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/middleware"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/Xavier-Tse/lunar-gate/utils/pwd"
)

type UpdatePasswordRequest struct {
	OldPwd string `json:"oldPwd" binding:"required"`
	Pwd    string `json:"pwd" binding:"required,max=64"`
	RePwd  string `json:"rePwd" binding:"required,max=64"`
}

func (UserApi) UpdatePassword(c *gin.Context) {
	var cr UpdatePasswordRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailBinding(err, c)
		return
	}
	claims := &middleware.GetAuth(c).ClaimsUserInfo

	var user model.User
	err := global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}

	if !pwd.Verify(user.Password, cr.OldPwd) {
		res.FailWithMessage("原密码错误", c)
		return
	}

	if cr.Pwd != cr.RePwd {
		res.FailWithMessage("两次密码不一致", c)
		return
	}
	hashPwd := pwd.Hash(cr.Pwd)
	global.DB.Model(&user).Update("password", hashPwd)
	res.OkWithMessage("修改密码成功", c)
}
