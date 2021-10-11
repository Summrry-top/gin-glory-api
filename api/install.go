package api

import (
	"github.com/Summrry-top/gin-glory-api/global"
	"github.com/Summrry-top/gin-glory-api/services"
	"github.com/gin-gonic/gin"
)

func PostInstall(c *gin.Context) {
	if global.ServerConfig.Install {
		Success(c, services.Err400("已安装"))
		return
	}
	mode := c.PostForm("mode")
	switch mode {
	case "checkDb":
		Success(c, services.CheckDb(c))
	case "checkRedis":
		Success(c, services.CheckRedis(c))
	case "install":
		Success(c, services.Install(c))
	default:
		Err400(c)
	}
}
