package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lunarise-dev/lunar-gate/common/res"
	"github.com/lunarise-dev/lunar-gate/service/redis_service"
	"github.com/lunarise-dev/lunar-gate/utils/jwts"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.Abort()
			res.FailWithMessage("请登陆", c)
			return
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
