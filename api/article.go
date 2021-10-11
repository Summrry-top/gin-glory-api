package api

import (
	"github.com/Summrry-top/gin-glory-api/models/orm"
	"github.com/Summrry-top/gin-glory-api/services"
	"github.com/Summrry-top/gin-glory-api/utils"
	"github.com/gin-gonic/gin"
)

func GetArticle(c *gin.Context) {
	g := services.GetParamDefault()
	_ = c.ShouldBind(&g)
	if !utils.VerifyParam(g) {
		Err400(c)
		return
	}
	if utils.Zero(g.Id) {
		Success(c, services.GetArticleAll(g))
		return
	}
	Success(c, services.GetArticleById(g.Id))

}

func PostArticle(c *gin.Context) {
	mode := c.GetString("mode")
	if mode == "delete" {
		id := c.PostForm("id")
		Success(c, services.PostArticleDelete(id))
		return
	}
	var o orm.Article
	if err := c.ShouldBind(&o); err != nil || o.Id < 0 {
		Err400(c)
		return
	}
	switch mode {
	case "add":
		Success(c, services.PostArticleAdd(&o))
	case "edit":
		Success(c, services.PostArticleEdit(&o))
	default:
		Err400(c)
	}
}
