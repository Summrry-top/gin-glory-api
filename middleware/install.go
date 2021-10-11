package middleware

import (
	"github.com/Summrry-top/gin-glory-api/global"
	"github.com/gin-gonic/gin"
)

func Install() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !global.ServerConfig.Install {
			c.JSON(200, "未初始化")
			c.Abort()
			return
		}
		c.Next()
	}
}
