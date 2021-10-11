package orm

import (
	"github.com/Summrry-top/gin-glory-api/models/response"
	"github.com/Summrry-top/gin-glory-api/utils"
)

// 文章分类
type ArticleSort struct {
	ArticleTag
	Alias       string `gorm:"comment:别名;" form:"alias"`
	Description string `gorm:"comment:描述;" form:"description"`
}

// 文章分类json数据
func (s *ArticleSort) GetArticleSortJson(expand interface{}) *response.ArticleSortJson {
	return &response.ArticleSortJson{
		Id:          s.Id,
		Name:        s.Name,
		Alias:       s.Alias,
		Description: s.Description,
		IsShow:      s.IsShow,
		Expand:      expand,
		CreateTime:  utils.FormatTime(s.CreatedAt),
		UpdateTime:  utils.FormatTime(s.UpdatedAt),
	}
}

// 获取分类id和名称
func (s *ArticleSort) GetArticleSortName() response.Name {
	return response.Name{
		Id:   s.Id,
		Name: s.Name,
	}
}
