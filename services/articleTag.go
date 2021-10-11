package services

import (
	"github.com/Summrry-top/gin-glory-api/models/orm"
	"github.com/Summrry-top/gin-glory-api/models/request"
	"github.com/Summrry-top/gin-glory-api/models/response"
	"github.com/Summrry-top/gin-glory-api/utils"
	"strconv"
)

// 内部 获取标签名称
func ArticleTagNameByIdArray(array []string) []response.Name {
	var articleTag orm.ArticleTag
	var articleTags []response.Name
	for _, v := range array {
		if !GetOne("id", v, &articleTag) {
			continue
		}
		articleTags = append(articleTags, articleTag.GetArticleTagName())
	}
	return articleTags
}

// 通过id获取一个文章标签 包含其中的全部文章
func GetArticleTagById(p request.GetParam) *response.Response {
	var articleTag orm.ArticleTag
	if !GetOne("id", p.Id, &articleTag) {
		return Err400("文章标签不存在！")
	}
	expand := GetArticleAllByField("article_tag_id LIKE ?", articleTag.Id, p)
	return Success("文章标签单条数据", articleTag.GetArticleTagJson(expand))
}

// 获取全部文章标签下的 包含其中的全部文章
func GetArticleTagAll(p request.GetParam) *response.Response {
	var articleTags []orm.ArticleTag
	var pageData response.PageData
	if !GetPagination(p, &pageData, &articleTags) {
		return Success("文章标签分页错误", pageData)
	}
	var articleTagsJson []*response.ArticleTagJson
	var article orm.Article
	for _, v := range articleTags {
		var expand response.ArticleTagsExpand
		expand.Count = GetCountByField(article, "article_tag_id LIKE ?", v.Id)
		articleTagsJson = append(articleTagsJson, v.GetArticleTagJson(expand))
	}
	pageData.Data = articleTagsJson
	return Success("文章标签多条数据", pageData)
}

// 新增一条文章标签
func PostArticleTagAdd(p *orm.ArticleTag) *response.Response {
	if data := PostArticleTagVerify(p); data != nil {
		return data
	}
	if CreateOne(p) {
		return Ok()
	}
	return Err400("文章标签已存在！新增失败！")

}

// 修改一条文章标签
func PostArticleTagEdit(p *orm.ArticleTag) *response.Response {
	if utils.Zero(p.Id) {
		return Err400("id错误")
	}
	if data := PostArticleTagVerify(p); data != nil {
		return data
	}
	if UpdateOne(p) {
		return Ok()
	}
	return Err400("文章标签更新失败")
}

// 删除一条文章标签
func PostArticleTagDelete(idStr string) *response.Response {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return Err400("id错误")
	}
	if DeleteOne(id, new(orm.ArticleTag)) {
		return Ok()
	}
	return Err400("文章标签删除失败")
}

// 校验文章标签
func PostArticleTagVerify(p *orm.ArticleTag) *response.Response {
	if !utils.VerifyName(p.Name) {
		return Err400("名称不合法")
	}
	return nil
}
