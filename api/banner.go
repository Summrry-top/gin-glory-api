package api

import (
	"github.com/Summrry-top/gin-glory-api/models/orm"
	"github.com/Summrry-top/gin-glory-api/services"
	"github.com/Summrry-top/gin-glory-api/utils"
	"github.com/gin-gonic/gin"
)

func GetBanner(c *gin.Context) {
	g := services.GetParamDefault()
	_ = c.ShouldBind(&g)
	if !utils.VerifyParam(g) {
		Err400(c)
		return
	}
	if utils.Zero(g.Id) {
		Success(c, services.GetBannerAll(g))
		return
	}
	Success(c, services.GetBannerById(g.Id))

}

func PostBanner(c *gin.Context) {
	mode := c.GetString("mode")
	if mode == "delete" {
		id := c.PostForm("id")
		Success(c, services.PostBannerDelete(id))
		return
	}
	var o orm.Banner
	if err := c.ShouldBind(&o); err != nil || o.Id < 0 {
		Err400(c)
		return
	}
	switch mode {
	case "add":
		Success(c, services.PostBannerAdd(&o))
	case "edit":
		Success(c, services.PostBannerEdit(&o))
	default:
		Err400(c)
	}
}
