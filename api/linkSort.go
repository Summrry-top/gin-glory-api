package api

import (
	"github.com/Summrry-top/gin-glory-api/models/orm"
	"github.com/Summrry-top/gin-glory-api/services"
	"github.com/Summrry-top/gin-glory-api/utils"
	"github.com/gin-gonic/gin"
)

func GetLinkSort(c *gin.Context) {
	g := services.GetParamDefault()
	_ = c.ShouldBind(&g)
	if !utils.VerifyParam(g) {
		Err400(c)
		return
	}
	if utils.Zero(g.Id) {
		Success(c, services.GetLinkSortAll(g))
		return
	}
	Success(c, services.GetLinkSortById(g))

}

func PostLinkSort(c *gin.Context) {
	mode := c.GetString("mode")
	id := c.PostForm("id")
	if mode == "delete" {
		Success(c, services.PostLinkSortDelete(id))
		return
	}
	var o orm.LinkSort
	if err := c.ShouldBind(&o); err != nil || o.Id < 0 {
		Err400(c)
		return
	}
	switch mode {
	case "add":
		Success(c, services.PostLinkSortAdd(&o))
	case "edit":
		Success(c, services.PostLinkSortEdit(&o))
	default:
		Err400(c)
	}
}
