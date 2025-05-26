package user_api

import (
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/service/redis_service"
	"github.com/Xavier-Tse/lunar-gate/utils/jwts"
	"github.com/gin-gonic/gin"
	"time"
)

func (UserApi) Logout(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	token := auth[7:]

	claims, _ := jwts.ParseToken(token)
	redis_service.AddToken(token, time.Duration(claims.ExpiresAt-time.Now().Unix())*time.Second)

	res.OkWithMessage("注销成功", c)
}
