package middleware

import (
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin := c.GetBool("isAdmin")
		if !isAdmin {
			c.JSON(200, "无权限")
			c.Abort()
			return
		}
		mode := c.GetString("mode")
		if mode == "" {
			c.JSON(200, "参数错误")
			c.Abort()
			return
		}
	}
}
