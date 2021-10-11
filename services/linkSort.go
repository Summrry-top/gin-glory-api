package services

import (
	"github.com/Summrry-top/gin-glory-api/models/orm"
	"github.com/Summrry-top/gin-glory-api/models/request"
	"github.com/Summrry-top/gin-glory-api/models/response"
	"github.com/Summrry-top/gin-glory-api/utils"
	"strconv"
)

// 内部 获取友链分类名称
func LinkSortNameByIdArray(array []string) []response.Name {
	var linkSort orm.LinkSort
	var linkSortNames []response.Name
	for _, v := range array {
		if ok := GetOne("id", v, &linkSort); !ok {
			continue
		}
		linkSortNames = append(linkSortNames, linkSort.GetLinkSortName())
	}
	return linkSortNames
}

// 通过id获取一个友链分类 包含其中的全部友链
func GetLinkSortById(p request.GetParam) *response.Response {
	var LinkSort orm.LinkSort
	if !GetOne("id", p.Id, &LinkSort) {
		return Err400("友链分类不存在！")
	}
	expand := GetLinkAllByField("link_sort_id LIKE ?", LinkSort.Id, p)
	return Success("友链分类单条数据", LinkSort.GetLinkSortJson(expand))
}

// 获取全部友链分类下的 包含其中的全部友链
func GetLinkSortAll(p request.GetParam) *response.Response {
	var LinkSorts []orm.LinkSort
	var pageData response.PageData
	if !GetPagination(p, &pageData, &LinkSorts) {
		return Fail("友链分类没有数据！", pageData)
	}
	var LinkSortsJson []*response.LinkSortJson
	var Link orm.Link
	for _, v := range LinkSorts {
		var expand response.LinkSortsExpand
		expand.Count = GetCountByField(Link, "Link_sort_id LIKE ?", v.Id)
		LinkSortsJson = append(LinkSortsJson, v.GetLinkSortJson(expand))
	}
	pageData.Data = LinkSortsJson
	return Success("友链分类多条数据", pageData)
}

// 新增一条友链分类
func PostLinkSortAdd(p *orm.LinkSort) *response.Response {
	if data := PostLinkSortVerify(p); data != nil {
		return data
	}
	if CreateOne(p) {
		return Ok()
	}
	return Err400("友链分类已存在！新增失败！")

}

// 修改一条友链分类
func PostLinkSortEdit(p *orm.LinkSort) *response.Response {
	if utils.Zero(p.Id) {
		return Err400("id错误")
	}
	if data := PostLinkSortVerify(p); data != nil {
		return data
	}
	if UpdateOne(p) {
		return Ok()
	}
	return Err400("友链分类更新失败")
}

// 删除一条友链分类
func PostLinkSortDelete(idStr string) *response.Response {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return Err400("id错误")
	}
	if DeleteOne(id, new(orm.LinkSort)) {
		return Ok()
	}
	return Err400("友链分类删除失败")
}

// 校验友链分类
func PostLinkSortVerify(p *orm.LinkSort) *response.Response {
	if !utils.VerifyName(p.Name) {
		return Err400("名称不合法")
	}
	if !utils.VerifyName(p.Description) {
		return Err400("描述不合法")
	}
	return nil
}
