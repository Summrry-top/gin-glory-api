package api

import (
	"github.com/Summrry-top/gin-glory-api/services"
	"github.com/gin-gonic/gin"
)

func PostVerifyCode(c *gin.Context) {
	email := c.PostForm("email")
	Success(c, services.CreateVerifyCode(email))
}
