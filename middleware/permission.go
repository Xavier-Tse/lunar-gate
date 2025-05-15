package middleware

import (
	"fmt"
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/utils/jwts"
	"github.com/gin-gonic/gin"
)

func Permission() gin.HandlerFunc {
	return func(c *gin.Context) {
		_claims, ok := c.Get("claims")
		if !ok {
			return
		}
		claims, _ := _claims.(*jwts.Claims)
		if claims.IsAdmin {
			c.Next()
			return
		}
		ok, _ = global.Casbin.Enforce(fmt.Sprintf("%d", claims.UserID), c.Request.URL.Path, c.Request.Method)
		if !ok {
			res.FailWithMessage("鉴权失败", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
