package middleware

import (
	"fmt"
	"github.com/Summrry-top/gin-glory-api/global"
	"github.com/Summrry-top/gin-glory-api/models/request"

	// "github.com/Summrry-top/gin-glory-api/models"
	"github.com/Summrry-top/gin-glory-api/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func MyJwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token
		var p request.Param
		if err := c.ShouldBind(&p); err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{
				"param": p,
				"error": err,
			})
			c.Abort()
			return
		}
		if p.LoginToken == "" {
			fmt.Println("未登录", p.Mode)
			c.Set("mode", p.Mode)
			c.Set("isAdmin", false)
			c.Next()
			return
		}
		//校验
		claims, err := utils.ParseToken(p.LoginToken)
		if err != nil {
			fmt.Println("验证失败！")
			global.Logger.Error("token错误",
				zap.String("ip", c.ClientIP()),
				zap.String("method", c.Request.Method),
				zap.String("path", c.Request.URL.Path),
				zap.String("query", c.Request.URL.RawQuery),
				zap.String("user-agent", c.Request.UserAgent()),
			)
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "token错误",
			})
			c.Abort()
			return
		}
		fmt.Println(claims, "验证成功")
		c.Set("mode", p.Mode)
		c.Set("account", claims.Account)
		c.Set("isAdmin", claims.IsAdmin)
		c.Next()
	}
}
