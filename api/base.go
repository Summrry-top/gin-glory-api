package api

import (
	"github.com/Summrry-top/gin-glory-api/global"
	"github.com/Summrry-top/gin-glory-api/services"
	"github.com/gin-gonic/gin"
)

func Err404(c *gin.Context) {
	data := services.Err404()
	c.JSON(global.OK, data)
}

func Err400(c *gin.Context) {
	data := services.Err400("参数错误！")
	c.JSON(global.OK, data)
}
func Err500(c *gin.Context) {
	data := services.Err500()
	c.JSON(global.OK, data)
}
func Success(c *gin.Context, data interface{}) {
	c.JSON(global.OK, data)
}
