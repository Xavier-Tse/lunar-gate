package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lunarise-dev/lunar-gate/common/res"
	"github.com/lunarise-dev/lunar-gate/global"
	"github.com/lunarise-dev/lunar-gate/model"
	"github.com/lunarise-dev/lunar-gate/utils/email"
	"github.com/lunarise-dev/lunar-gate/utils/pwd"
)

type RegisterRequest struct {
	EmailID    string `json:"emailID" binding:"required"`
	EmailCode  string `json:"emailCode" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"rePassword" binding:"required"`
}

func (UserApi) Register(c *gin.Context) {
	var cr RegisterRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailBinding(err, c)
		return
	}

	if !email.Verify(cr.EmailID, cr.Email, cr.EmailCode) {
		res.FailWithMessage("邮箱验证失败", c)
		email.Remove(cr.EmailID)
		return
	}
	if cr.Password != cr.RePassword {
		res.FailWithMessage("两次密码不一致", c)
		return
	}

	var user model.User
	err := global.DB.Take(&user, "email = ?", cr.Email).Error
	if err == nil {
		res.FailWithMessage("该邮箱已被使用", c)
		email.Remove(cr.EmailID)
		return
	}

	hashPwd := pwd.Hash(cr.Password)
	fmt.Println(len(cr.Email))
	err = global.DB.Create(&model.User{
		Username: cr.Email,
		Nickname: "邮箱用户",
		Email:    cr.Email,
		Password: hashPwd,
	}).Error
	if err != nil {
		res.FailWithMessage("注册失败", c)
		email.Remove(cr.EmailID)
		return
	}
	res.OkWithMessage("用户注册成功", c)
}
