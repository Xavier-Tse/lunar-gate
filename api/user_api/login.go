package user_api

import (
	"github.com/gin-gonic/gin"
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/Xavier-Tse/lunar-gate/utils/captcha"
	"github.com/Xavier-Tse/lunar-gate/utils/jwts"
	"github.com/Xavier-Tse/lunar-gate/utils/pwd"
	"github.com/sirupsen/logrus"
)

type LoginRequest struct {
	Username    string `json:"username" binding:"required" label:"用户名"`
	Password    string `json:"password" binding:"required" label:"密码"`
	CaptchaID   string `json:"captchaID"`
	CaptchaCode string `json:"captchaCode"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (UserApi) Login(c *gin.Context) {
	var cr LoginRequest
	if err := c.ShouldBind(&cr); err != nil {
		res.FailBinding(err, c)
		return
	}

	if global.Config.Captcha.Enable {
		if cr.CaptchaID == "" || cr.CaptchaCode == "" {
			res.FailWithMessage("请输入验证码内容", c)
			return
		}
		if !captcha.Store.Verify(cr.CaptchaID, cr.CaptchaCode, true) {
			res.FailWithMessage("验证失败", c)
			return
		}
	}

	var user model.User
	if err := global.DB.Preload("RoleList").Take(&user, "username = ?", cr.Username).Error; err != nil {
		res.FailWithMessage("用户名或密码错误", c)
		return
	}

	if !pwd.Verify(user.Password, cr.Password) {
		res.FailWithMessage("用户名或密码错误", c)
		return
	}

	var roleList []uint
	for _, u := range user.RoleList {
		roleList = append(roleList, u.ID)
	}

	token, err := jwts.GenerateToken(jwts.ClaimsUserInfo{
		UserID:   user.ID,
		Username: user.Username,
		RoleList: roleList,
	})
	if err != nil {
		logrus.Errorf("jwt生成token失败 %v", err)
		res.FailWithMessage("用户登陆失败", c)
		return
	}

	res.OkWithData(LoginResponse{
		Token: token,
	}, c)
}
