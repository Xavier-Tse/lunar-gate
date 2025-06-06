package user_api

import (
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/core"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/Xavier-Tse/lunar-gate/utils/captcha"
	"github.com/Xavier-Tse/lunar-gate/utils/jwts"
	"github.com/Xavier-Tse/lunar-gate/utils/pwd"
	"github.com/gin-gonic/gin"
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

	if global.Config.Info.Login.Captcha.Enable {
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
		IsAdmin:  user.IsAdmin,
	})
	if err != nil {
		logrus.Errorf("jwt生成token失败 %v", err)
		res.FailWithMessage("用户登陆失败", c)
		return
	}

	ip := c.ClientIP()
	addr := core.GetIpAddr(ip)
	ua := c.Request.UserAgent()
	global.DB.Create(&model.UserLogin{
		UserID: user.ID,
		IP:     ip,
		Addr:   addr,
		UA:     ua,
	})

	logrus.Infof("用户 %s 登陆成功, IP: %s, 地址: %s, UA: %s", user.Username, ip, addr, ua)

	res.OkWithData(LoginResponse{
		Token: token,
	}, c)
}
