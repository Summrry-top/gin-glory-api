package orm

import (
	"github.com/Summrry-top/gin-glory-api/models/response"
	"github.com/Summrry-top/gin-glory-api/utils"
)

type ArticleTag struct {
	Model
	Name   string `gorm:"comment:名称;" form:"name"`
	IsShow bool   `gorm:"comment:显示;" form:"is_show"`
}

func (t ArticleTag) GetArticleTagJson(expand interface{}) *response.ArticleTagJson {
	return &response.ArticleTagJson{
		Id:         t.Id,
		Name:       t.Name,
		IsShow:     t.IsShow,
		Expand:     expand,
		CreateTime: utils.FormatTime(t.CreatedAt),
		UpdateTime: utils.FormatTime(t.UpdatedAt),
	}
}

func (t ArticleTag) GetArticleTagName() response.Name {
	return response.Name{
		Id:   t.Id,
		Name: t.Name,
	}
}
