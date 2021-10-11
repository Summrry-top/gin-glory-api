package services

import (
	"fmt"
	"github.com/Summrry-top/gin-glory-api/models/orm"
	"github.com/Summrry-top/gin-glory-api/models/request"
	"github.com/Summrry-top/gin-glory-api/models/response"
	"github.com/Summrry-top/gin-glory-api/utils"
	"strconv"
)

// 通过id获取一条配置信息数据
func GetOptionById(id int) *response.Response {
	var option orm.Option
	if ok := GetOne("id", id, &option); !ok {
		return Err400("配置信息不存在！")
	}
	return Success("配置信息单条数据", option.GetOptionJson())
}

// 获取多条配置信息数据
func GetOptionAll(p request.GetParam) *response.Response {
	var options []orm.Option
	var pageData response.PageData
	if ok := GetPagination(p, &pageData, &options); !ok {
		return Fail("配置信息没有数据！", pageData)
	}
	var optionsJson []*response.OptionJson
	for _, v := range options {
		optionsJson = append(optionsJson, v.GetOptionJson())
	}
	pageData.Data = optionsJson
	return Success("配置信息多条数据", pageData)
}

// 新增一条配置信息数据
func PostOptionAdd(p *orm.Option) *response.Response {
	if data := PostOptionVerify(p); data != nil {
		return data
	}
	if CreateOne(p) {
		return Ok()
	}
	fmt.Println(p)
	return Err400("配置信息已存在！新增失败！")

}

// 修改一条配置信息数据
func PostOptionEdit(p *orm.Option) *response.Response {
	if utils.Zero(p.Id) {
		return Err400("id错误")
	}
	if data := PostOptionVerify(p); data != nil {
		return data
	}
	if UpdateOne(p) {
		return Ok()
	}
	return Err400("配置信息更新失败")
}

// 删除一条配置信息数据
func PostOptionDelete(idStr string) *response.Response {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return Err400("id错误")
	}
	if DeleteOne(id, new(orm.Option)) {
		return Ok()
	}
	return Err400("配置信息删除失败")
}

// 校验配置信息数据
func PostOptionVerify(p *orm.Option) *response.Response {
	if utils.VerifyName(p.Title) {
		return Err400("标题不合法")
	}
	if utils.VerifyName(p.Copy) {
		return Err400("copy不合法")
	}
	if utils.VerifyName(p.Description) {
		return Err400("描述不合法")
	}
	if utils.VerifyUrl(p.SiteUrl) {
		return Err400("url不合法")
	}
	if utils.VerifyUrl(p.SiteImg) {
		return Err400("img不合法")
	}
	if utils.VerifyUrl(p.SiteIco) {
		return Err400("img不合法")
	}

	return nil
}
