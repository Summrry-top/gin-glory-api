package services

import (
	"github.com/Summrry-top/gin-glory-api/models/orm"
	"github.com/Summrry-top/gin-glory-api/models/request"
	"github.com/Summrry-top/gin-glory-api/models/response"
	"github.com/Summrry-top/gin-glory-api/utils"
	"strconv"
)

// 通过id获取一条轮播图数据
func GetBannerById(id int) *response.Response {
	var banner orm.Banner
	if !GetOne("id", id, &banner) {
		return Err400("Banner不存在！")
	}
	return Success("Banner单条数据", banner.GetBannerJson())
}

// 获取多条轮播图数据
func GetBannerAll(p request.GetParam) *response.Response {
	var banners []orm.Banner
	var pageData response.PageData
	if !GetPagination(p, &pageData, &banners) {
		return Fail("轮播图没有数据", pageData)
	}
	var bannersJson []*response.BannerJson
	for _, v := range banners {
		bannersJson = append(bannersJson, v.GetBannerJson())
	}
	pageData.Data = bannersJson
	return Success("Banner多条数据", pageData)
}

// 新增一条轮播图数据
func PostBannerAdd(p *orm.Banner) *response.Response {
	if data := PostBannerVerify(p); data != nil {
		return data
	}
	if CreateOne(p) {
		return Ok()
	}
	return Err400("已存在！新增失败！")

}

// 修改一条轮播图数据
func PostBannerEdit(p *orm.Banner) *response.Response {
	if utils.Zero(p.Id) {
		return Err400("id错误")
	}
	if data := PostBannerVerify(p); data != nil {
		return data
	}
	if UpdateOne(p) {
		return Ok()
	}
	return Err400("更新失败")
}

// 删除一条轮播图数据
func PostBannerDelete(idStr string) *response.Response {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return Err400("id错误")
	}
	if DeleteOne(id, new(orm.Banner)) {
		return Ok()
	}
	return Err400("删除失败")
}

// 校验轮播图数据
func PostBannerVerify(p *orm.Banner) *response.Response {
	if !utils.VerifyUrl(p.Url) {
		return Err400("url不合法")
	}
	if !utils.VerifyUrl(p.Img) {
		return Err400("img不合法")
	}
	if !utils.VerifyName(p.Title) {
		return Err400("标题不合法")
	}
	if !utils.VerifyName(p.Description) {
		return Err400("描述不合法")
	}
	return nil
}
