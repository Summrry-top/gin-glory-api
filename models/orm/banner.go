package orm

import (
	"github.com/Summrry-top/gin-glory-api/models/response"
	"github.com/Summrry-top/gin-glory-api/utils"
)

type Banner struct {
	Model
	Title       string `gorm:"comment:标题;unique" form:"title"`
	Url         string `gorm:"comment:地址;" form:"url"`
	Img         string `gorm:"comment:图片;" form:"img"`
	Description string `gorm:"comment:描述;" form:"description"`
}

func (b *Banner) GetBannerJson() *response.BannerJson {
	return &response.BannerJson{
		Id:          b.Id,
		Title:       b.Title,
		Description: b.Description,
		Url:         b.Url,
		Img:         b.Img,
		CreateTime:  utils.FormatTime(b.CreatedAt),
		UpdateTime:  utils.FormatTime(b.UpdatedAt),
	}
}
