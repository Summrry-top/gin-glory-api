package services

import (
	"github.com/Summrry-top/gin-glory-api/global"
	"github.com/Summrry-top/gin-glory-api/models/request"
	"github.com/Summrry-top/gin-glory-api/models/response"
	"github.com/Summrry-top/gin-glory-api/utils"
)

// 分页信息处理
func PageDataDefault(p *response.PageData, g request.GetParam) {
	p.CurrentPage = g.Page
	p.PageSize = g.Limit
	p.TotalPage = utils.TotalPage(p.Count, int64(p.PageSize))
}
func Err404() *response.Response {
	return &response.Response{
		Code: 404,
		Msg:  "资源不存在！",
		Data: global.Nil,
	}
}
func Err400(msg string) *response.Response {
	return &response.Response{
		Code: 400,
		Msg:  msg,
		Data: global.Nil,
	}
}

func Err500() *response.Response {
	return &response.Response{
		Code: 500,
		Msg:  "服务器错误！",
		Data: global.Nil,
	}
}
func Success(msg string, data interface{}) *response.Response {
	return &response.Response{
		Code: 200,
		Msg:  msg,
		Data: data,
	}
}

// 分页错误使用
func Fail(msg string, data interface{}) *response.Response {
	return &response.Response{
		Code: 400,
		Msg:  "分页错误",
		Data: data,
	}
}
func Ok() *response.Response {
	return &response.Response{
		Code: 200,
		Msg:  "ok",
		Data: global.Nil,
	}
}
