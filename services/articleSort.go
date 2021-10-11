package services

import (
	"github.com/Summrry-top/gin-glory-api/models/orm"
	"github.com/Summrry-top/gin-glory-api/models/request"
	"github.com/Summrry-top/gin-glory-api/models/response"
	"github.com/Summrry-top/gin-glory-api/utils"
	"strconv"
)

// 内部 获取分类名称
func ArticleSortNameByIdArray(array []string) []response.Name {
	var articleSort orm.ArticleSort
	var articleSorts []response.Name
	for _, v := range array {
		if !GetOne("id", v, &articleSort) {
			continue
		}
		articleSorts = append(articleSorts, articleSort.GetArticleSortName())
	}
	return articleSorts
}

// 通过id获取一个文章分类 包含其中的全部文章
func GetArticleSortById(p request.GetParam) *response.Response {
	var articleSort orm.ArticleSort
	if !GetOne("id", p.Id, &articleSort) {
		return Err400("文章分类不存在！")
	}
	expand := GetArticleAllByField("article_sort_id LIKE ?", articleSort.Id, p)
	return Success("文章分类单条数据", articleSort.GetArticleSortJson(expand))
}

// 获取全部文章分类下的 包含其中的全部文章
func GetArticleSortAll(p request.GetParam) *response.Response {
	var articleSorts []orm.ArticleSort
	var pageData response.PageData
	if !GetPagination(p, &pageData, &articleSorts) {
		return Err400("文章分类没有数据！")
	}
	var articleSortsJson []*response.ArticleSortJson
	var article orm.Article
	for _, v := range articleSorts {
		var expand response.ArticleSortsExpand
		expand.Count = GetCountByField(article, "article_sort_id LIKE ?", v.Id)
		articleSortsJson = append(articleSortsJson, v.GetArticleSortJson(expand))
	}
	pageData.Data = articleSortsJson
	return Success("文章分类多条数据", pageData)
}

// 新增一条文章分类
func PostArticleSortAdd(p *orm.ArticleSort) *response.Response {
	if data := PostArticleSortVerify(p); data != nil {
		return data
	}
	if CreateOne(p) {
		return Ok()
	}
	return Err400("文章分类已存在！新增失败！")

}

// 修改一条文章分类
func PostArticleSortEdit(p *orm.ArticleSort) *response.Response {
	if utils.Zero(p.Id) {
		return Err400("id错误")
	}
	if data := PostArticleSortVerify(p); data != nil {
		return data
	}
	if UpdateOne(p) {
		return Ok()
	}
	return Err400("文章分类更新失败")
}

// 删除一条文章分类
func PostArticleSortDelete(idStr string) *response.Response {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return Err400("id错误")
	}
	if DeleteOne(id, new(orm.ArticleSort)) {
		return Ok()
	}
	return Err400("文章分类删除失败")
}

// 校验文章分类
func PostArticleSortVerify(p *orm.ArticleSort) *response.Response {
	if !utils.VerifyName(p.Name) {
		return Err400("名称不合法")
	}
	if !utils.VerifyName(p.Description) {
		return Err400("描述不合法")
	}
	return nil
}
