package api

import (
	"github.com/Summrry-top/gin-glory-api/models/orm"
	"github.com/Summrry-top/gin-glory-api/services"
	"github.com/Summrry-top/gin-glory-api/utils"
	"github.com/gin-gonic/gin"
)

func GetLink(c *gin.Context) {
	g := services.GetParamDefault()
	_ = c.ShouldBind(&g)
	if !utils.VerifyParam(g) {
		Err400(c)
		return
	}
	if utils.Zero(g.Id) {
		Success(c, services.GetLinkAll(g))
		return
	}
	Success(c, services.GetLinkById(g.Id))

}

func PostLink(c *gin.Context) {
	mode := c.GetString("mode")
	id := c.PostForm("id")
	if mode == "delete" {
		Success(c, services.PostLinkDelete(id))
		return
	}
	var o orm.Link
	if err := c.ShouldBind(&o); err != nil || o.Id < 0 {
		Err400(c)
		return
	}
	switch mode {
	case "add":
		Success(c, services.PostLinkAdd(&o))
	case "edit":
		Success(c, services.PostLinkEdit(&o))
	default:
		Err400(c)
	}
}
