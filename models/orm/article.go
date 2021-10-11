package orm

import (
	"github.com/Summrry-top/gin-glory-api/models/response"
)

// 文章orm
type Article struct {
	Model
	UserId        int    `gorm:"comment:作者"`
	Title         string `gorm:"comment:标题;" form:"title"`
	Description   string `gorm:"comment:描述;" form:"description"`
	Content       string `gorm:"comment:内容;" form:"content"`
	ArticleSortId string `gorm:"comment:分类;" form:"article_sort_id"`
	ArticleTagId  string `gorm:"comment:标签;" form:"article_tag_id"`
	Top           int    `gorm:"comment:置顶;" form:"top"`
	IsShow        bool   `gorm:"comment:显示;" form:"is_show"`
	ImgSrc        string `gorm:"comment:封面;" form:"img_src"`
	FontCount     int    `gorm:"comment:字数;" form:"font_count"`
}

// 一条文章json数据 文章详情页
func (a Article) GetArticleJson(format response.ArticleFormat) *response.ArticleJson {
	return &response.ArticleJson{
		Id:            a.Id,
		Title:         a.Title,
		Description:   a.Description,
		Content:       a.Content,
		Top:           a.Top,
		IsShow:        a.IsShow,
		Views:         0,
		FontCount:     a.FontCount,
		ArticleSortId: format.Sort,
		ArticleTagId:  format.Tag,
		UserId:        a.UserId,
		Expand:        format.Expand,
		CreateTime:    format.CreateTime,
		UpdateTime:    format.UpdateTime,
	}
}

// 多条文章json数据 不包含文章内容
func (a Article) GetArticlesJson(format response.ArticleFormat) *response.ArticlesJson {
	return &response.ArticlesJson{
		Id:            a.Id,
		Title:         a.Title,
		Description:   a.Description,
		Top:           a.Top,
		IsShow:        a.IsShow,
		ImgSrc:        a.ImgSrc,
		Views:         0,
		FontCount:     a.FontCount,
		ArticleSortId: format.Sort,
		ArticleTagId:  format.Tag,
		UserId:        a.UserId,
		Expand:        format.Expand,
		CreateTime:    format.CreateTime,
		UpdateTime:    format.CreateTime,
	}
}
