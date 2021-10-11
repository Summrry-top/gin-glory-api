package services

import (
	"github.com/Summrry-top/gin-glory-api/global"
	"github.com/Summrry-top/gin-glory-api/models/request"
)

func GetParamDefault() request.GetParam {
	return request.GetParam{
		Page:  global.Page,
		Limit: global.Limit,
		Order: global.Order,
		Cache: global.Cache,
	}
}
