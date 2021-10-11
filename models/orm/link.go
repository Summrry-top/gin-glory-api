package orm

import (
	"github.com/Summrry-top/gin-glory-api/models/response"
)

type Link struct {
	ArticleTag
	Url         string `gorm:"comment:地址;" form:"url"`
	HeadImg     string `gorm:"comment:头像;" form:"head_img"`
	LinkSortId  string `gorm:"comment:排序;" form:"sort"`
	Description string `gorm:"comment:描述;" form:"description"`
}

func (l *Link) GetLinkJson(format response.LinkFormat) *response.LinkJson {
	return &response.LinkJson{
		Id:          l.Id,
		Name:        l.Name,
		Url:         l.Url,
		HeadImg:     l.HeadImg,
		Description: l.Description,
		LinkSort:    format.Sort,
		IsShow:      l.IsShow,
		Expand:      format.Expand,
		CreateTime:  format.CreateTime,
		UpdateTime:  format.UpdateTime,
	}
}
