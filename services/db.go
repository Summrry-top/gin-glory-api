package services

import (
	"errors"
	"fmt"
	"github.com/Summrry-top/gin-glory-api/global"
	"github.com/Summrry-top/gin-glory-api/models/request"
	"github.com/Summrry-top/gin-glory-api/models/response"
	"github.com/Summrry-top/gin-glory-api/utils"
	"gorm.io/gorm"
)

// 获取一条数据
func GetOne(field string, value, dest interface{}) bool {
	err := global.Db.Where(field+"= ?", value).First(dest).Error
	if utils.Nil(err) {
		return true
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	panic(err)
}

// 添加一条数据
func CreateOne(data interface{}) bool {
	result := global.Db.Create(data)
	if utils.Nil(result.Error) {
		return true
	}
	fmt.Println(result.Error)
	return false
}

// 更新一条数据
func UpdateOne(data interface{}) bool {
	fmt.Println(data)
	result := global.Db.Model(data).Updates(data)
	fmt.Println(result.Error, result.RowsAffected)
	if utils.Nil(result.Error) && result.RowsAffected == 1 {
		return true
	}
	return false
}

// 删除一条数据
func DeleteOne(id int, dest interface{}) bool {
	result := global.Db.Delete(dest, id)
	fmt.Println(result.Error, result.RowsAffected)
	if utils.Nil(result.Error) && result.RowsAffected == 1 {
		return true
	}
	return false
}

// 获取分页数据
// 请求参数 分页数据盒子 分页具体数据
// 是否有错误 服务结束响应数据
func GetPagination(p request.GetParam, pageData *response.PageData, dest interface{}) bool {
	err := global.Db.Model(dest).Count(&pageData.Count).Error
	if utils.NotNil(err) {
		panic(err)
	}
	if utils.Zero(int(pageData.Count)) {
		return false
	}
	PageDataDefault(pageData, p)
	if int64(pageData.CurrentPage) > pageData.TotalPage {
		return false
	}
	err = global.Db.Order(p.Order).Limit(p.Limit).Offset((p.Page - 1) * p.Limit).Find(dest).Error
	if utils.NotNil(err) {
		panic(err)
	}

	return true
}

// 获取分类/标签下文章分页
func GetPaginationByField(query string, id int, p request.GetParam, pageData *response.PageData, dest interface{}) bool {
	tx := global.Db.Model(dest).Where(query, utils.Join(id))
	if utils.NotNil(tx.Error) {
		panic(tx.Error)
	}
	tx.Count(&pageData.Count)
	PageDataDefault(pageData, p)
	if utils.Zero(int(pageData.Count)) {
		return false
	}

	if int64(pageData.CurrentPage) > pageData.TotalPage {
		return false
	}
	tx.Order(p.Order).Limit(p.Limit).Offset((p.Page - 1) * p.Limit).Find(dest)
	if utils.NotNil(tx.Error) {
		panic(tx.Error)
	}
	return true
}

func GetCountByField(model interface{}, query string, id int) int64 {
	var count int64
	global.Db.Model(model).Where(query, utils.Join(id)).Count(&count)
	return count
}
