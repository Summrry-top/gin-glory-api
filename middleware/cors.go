package middleware

import (
	"github.com/Summrry-top/gin-glory-api/global"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Credentials", global.AllowCredentials)
		c.Header("Access-Control-Allow-Origin", c.GetHeader("Origin")) // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Methods", global.AllowMethods)
		c.Header("Access-Control-Allow-Headers", global.AllowHeaders)
		c.Header("Access-Control-Expose-Headers", global.ExposeHeaders)
		c.Next()
	}
}
