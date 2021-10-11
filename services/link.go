package services

import (
	"fmt"
	"github.com/Summrry-top/gin-glory-api/models/orm"
	"github.com/Summrry-top/gin-glory-api/models/request"
	"github.com/Summrry-top/gin-glory-api/models/response"
	"github.com/Summrry-top/gin-glory-api/utils"
	"strconv"
)

// 内部 友链格式化
func GetLinkFormat(link *orm.Link) response.LinkFormat {
	var linkFormat response.LinkFormat
	// 分类
	linkFormat.Sort = utils.ToArray(link.LinkSortId)
	// 拓展 分类名称
	linkFormat.Expand = LinkSortNameByIdArray(utils.ToArray(link.LinkSortId))
	// 时间
	linkFormat.CreateTime = utils.FormatTime(link.CreatedAt)
	linkFormat.UpdateTime = utils.FormatTime(link.UpdatedAt)
	return linkFormat
}

// 通过id获取一条友链数据
func GetLinkById(id int) *response.Response {
	var link orm.Link
	if !GetOne("id", id, &link) {
		return Err400("友链不存在！")
	}
	return Success("友链单条数据", link.GetLinkJson(GetLinkFormat(&link)))
}

// 获取全部友链数据
func GetLinkAll(p request.GetParam) *response.Response {
	var links []orm.Link
	var pageData response.PageData
	if !GetPagination(p, &pageData, &links) {
		return Fail("友链没有数据！", pageData)
	}
	var linksJson []*response.LinkJson
	for _, v := range links {
		linksJson = append(linksJson, v.GetLinkJson(GetLinkFormat(&v)))
	}
	pageData.Data = linksJson
	return Success("友链多条数据", pageData)
}

// 内部 获取分类下的全部友链
func GetLinkAllByField(query string, FieldId int, p request.GetParam) response.PageData {
	var links []orm.Link
	var pageData response.PageData
	if !GetPaginationByField(query, FieldId, p, &pageData, &links) {
		return pageData
	}

	var linksJson []*response.LinkJson
	for _, v := range links {
		linksJson = append(linksJson, v.GetLinkJson(GetLinkFormat(&v)))
	}
	fmt.Println(123)
	pageData.Data = linksJson
	return pageData
}

// 新增一条友链数据
func PostLinkAdd(p *orm.Link) *response.Response {
	if data := PostLinkVerify(p); data != nil {
		return data
	}
	if CreateOne(p) {
		return Ok()
	}
	return Err400("友链已存在！新增失败！")

}

// 修改一条友链数据
func PostLinkEdit(p *orm.Link) *response.Response {
	if utils.Zero(p.Id) {
		return Err400("id错误")
	}
	if data := PostLinkVerify(p); data != nil {
		return data
	}
	if UpdateOne(p) {
		return Ok()
	}
	return Err400("友链更新失败")
}

// 删除一条友链数据
func PostLinkDelete(idStr string) *response.Response {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return Err400("id错误")
	}
	if DeleteOne(id, new(orm.Link)) {
		return Ok()
	}
	return Err400("友链删除失败")
}

// 校验友链数据
func PostLinkVerify(p *orm.Link) *response.Response {
	//if utils.VerifyUrl(p.Url) {
	//	return Err400("url不合法")
	//}
	//if utils.VerifyUrl(p.Img) {
	//	return Err400("img不合法")
	//}
	if !utils.VerifyName(p.Name) {
		return Err400("名称不合法")
	}
	if !utils.VerifyName(p.Description) {
		return Err400("描述不合法")
	}
	return nil
}

//// 获取分类下的友链
//func GetLinkBySort(id int) *response.LinkSortExpand{
//	var links []orm.Link
//	var linkSortExpand  response.LinkSortExpand
//	global.Db.Where("link_sort_id LIKE ?",utils.Join(id)).Find(&links)
//	for _,v:=range links{
//		linkSortExpand.Links=append(linkSortExpand.Links,v.GetLinksJson(GetLinkFormat(&v)))
//	}
//	linkSortExpand.Count=len(links)
//	return &linkSortExpand
//}
