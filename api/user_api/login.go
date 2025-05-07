package user_api

import (
	"github.com/gin-gonic/gin"
	"github.com/lunarise-dev/lunar-gate/global"
	"github.com/lunarise-dev/lunar-gate/model"
	"github.com/lunarise-dev/lunar-gate/utils/jwts"
	"github.com/lunarise-dev/lunar-gate/utils/pwd"
	"github.com/sirupsen/logrus"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (UserApi) Login(c *gin.Context) {
	var cr LoginRequest
	if err := c.ShouldBind(&cr); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	var user model.User
	if err := global.DB.Preload("RoleList").Take(&user, "username = ?", cr.Username).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	if !pwd.Verify(user.Password, cr.Password) {
		c.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"message": "用户名或密码错误",
			"data":    nil,
		})
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
		c.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"message": "用户登陆失败",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "登陆成功",
		"data":    LoginResponse{Token: token},
	})
}
