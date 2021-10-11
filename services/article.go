package services

import (
	"github.com/Summrry-top/gin-glory-api/models/orm"
	"github.com/Summrry-top/gin-glory-api/models/request"
	"github.com/Summrry-top/gin-glory-api/models/response"
	"github.com/Summrry-top/gin-glory-api/utils"
	"strconv"
)

// 内部 文章格式化
func GetArticleFormat(article *orm.Article) response.ArticleFormat {
	var articleFormat response.ArticleFormat
	// 分类
	articleFormat.Sort = utils.ToArray(article.ArticleSortId)
	// 标签
	articleFormat.Tag = utils.ToArray(article.ArticleTagId)
	// 拓展 分类
	articleFormat.Expand.ArticleSort = ArticleSortNameByIdArray(articleFormat.Sort)
	// 拓展 标签
	articleFormat.Expand.ArticleTag = ArticleTagNameByIdArray(articleFormat.Tag)
	// 时间
	articleFormat.CreateTime = utils.FormatTime(article.CreatedAt)
	articleFormat.UpdateTime = utils.FormatTime(article.UpdatedAt)
	return articleFormat
}

// 通过id获取一条文章数据
func GetArticleById(id int) *response.Response {
	var article orm.Article
	if !GetOne("id", id, &article) {
		return Err400("文章不存在！")
	}
	return Success("文章单条数据", article.GetArticleJson(GetArticleFormat(&article)))
}

// 获取全部文章数据
func GetArticleAll(p request.GetParam) *response.Response {
	var articles []orm.Article
	var pageData response.PageData
	if !GetPagination(p, &pageData, &articles) {
		return Fail("文章没有数据", pageData)
	}
	var articlesJson []*response.ArticlesJson
	for _, v := range articles {
		articlesJson = append(articlesJson, v.GetArticlesJson(GetArticleFormat(&v)))
	}
	pageData.Data = articlesJson
	return Success("文章多条数据", pageData)
}

// 获取分类|标签下的全部文章
func GetArticleAllByField(query string, FieldId int, p request.GetParam) response.PageData {
	var articles []orm.Article
	var pageData response.PageData
	if !GetPaginationByField(query, FieldId, p, &pageData, &articles) {
		return pageData
	}
	var articlesJson []*response.ArticlesJson
	for _, v := range articles {
		articlesJson = append(articlesJson, v.GetArticlesJson(GetArticleFormat(&v)))
	}
	pageData.Data = articlesJson
	return pageData
}

// 新增一条文章数据
func PostArticleAdd(p *orm.Article) *response.Response {
	if data := PostArticleVerify(p); data != nil {
		return data
	}
	if CreateOne(p) {
		return Ok()
	}
	return Err400("文章已存在！新增失败！")

}

// 修改一条文章数据
func PostArticleEdit(p *orm.Article) *response.Response {
	if utils.Zero(p.Id) {
		return Err400("id错误")
	}
	if data := PostArticleVerify(p); data != nil {
		return data
	}
	if UpdateOne(p) {
		return Ok()
	}
	return Err400("文章更新失败")
}

// 删除一条文章数据
func PostArticleDelete(idStr string) *response.Response {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return Err400("id错误")
	}
	if DeleteOne(id, new(orm.Article)) {
		return Ok()
	}
	return Err400("文章删除失败")
}

// 校验文章数据
func PostArticleVerify(p *orm.Article) *response.Response {
	//if !utils.VerifyUrl(p.Url) {
	//	return Err400("url不合法")
	//}
	//if !utils.VerifyUrl(p.Img) {
	//	return Err400("img不合法")
	//}
	if !utils.VerifyName(p.Title) {
		return Err400("标题不合法")
	}
	if !utils.VerifyName(p.Description) {
		return Err400("描述不合法")
	}
	return nil
}

//// 获取分类下的文章
//func GetArticleBySort(id int) *response.ArticleSortExpand{
//	var articles []orm.Article
//	var articleSortExpand  response.ArticleSortExpand
//	global.Db.Where("article_sort_id LIKE ?",utils.Join(id)).Find(&articles)
//	for _,v:=range articles{
//		articleSortExpand.Articles=append(articleSortExpand.Articles,v.GetArticlesJson(GetArticleFormat(&v)))
//	}
//	articleSortExpand.Count=len(articles)
//	return &articleSortExpand
//}
