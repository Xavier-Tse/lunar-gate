package middleware

import (
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/service/redis_service"
	"github.com/Xavier-Tse/lunar-gate/utils"
	"github.com/Xavier-Tse/lunar-gate/utils/jwts"
	"github.com/gin-gonic/gin"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if utils.InList(global.Config.Router.WhiteRouter, c.Request.URL.Path) {
			c.Next()
			return
		}

		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			authHeader = c.Query("Authorization")
			if authHeader == "" {
				c.Abort()
				res.FailWithMessage("请登陆", c)
				return
			}
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.Abort()
			res.FailWithMessage("请求头格式错误", c)
			return
		}

		tokenString := authHeader[7:]
		if tokenString == "" {
			c.Abort()
			res.FailWithMessage("请求头格式错误", c)
			return
		}

		if redis_service.TokenBlack(tokenString) {
			c.Abort()
			res.FailWithMessage("用户已注销", c)
			return
		}

		claims, err := jwts.ParseToken(tokenString)
		if err != nil {
			c.Abort()
			res.FailWithMessage("token解析失败", c)
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

func GetAuth(c *gin.Context) *jwts.Claims {
	return c.MustGet("claims").(*jwts.Claims)
}
