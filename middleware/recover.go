package middleware

import (
	"fmt"
	"github.com/Summrry-top/gin-glory-api/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http/httputil"
)

func Recovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			// 打印错误栈信息
			fmt.Println(err)
			//
			c.JSON(200, gin.H{
				"code": 500,
				"msg":  "服务器错误！",
				"data": nil,
			})
			c.Abort()
			return
		}
	}()
	c.Next()
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
// stack 开启
func GinRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				// 检查断开的连接，因为它不是一个真正需要panic堆栈跟踪的条件。
				// if ne, ok := err.(*net.OpError); ok {
				// 	if se, ok := ne.Err.(*os.SyscallError); ok {
				// 		if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
				// 			// 连接断开
				// 			global.GvaLogger.Error(c.Request.URL.Path,
				// 				zap.Any("error", err),
				// 				zap.String("request", string(httpRequest)),
				// 			)
				// 			// 如果连接断开，我们就无法向它写入状态。
				// 			c.Error(err.(error)) // nolint: errcheck
				// 			c.Abort()
				// 			return
				// 		}
				// 	}
				// }

				// 写入错误日志
				global.Logger.Error("[Recovery from panic]",
					zap.Any("error", err),
					zap.String("request", string(httpRequest)),
				)
				log.Println(err)
				c.JSON(200, "500")
				c.Abort()
			}
		}()
		c.Next()
	}
}
