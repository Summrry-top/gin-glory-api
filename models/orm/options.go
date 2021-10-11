package orm

import (
	"github.com/Summrry-top/gin-glory-api/models/response"
	"github.com/Summrry-top/gin-glory-api/utils"
)

type Option struct {
	Model
	Title       string `gorm:"comment:标题;" form:"title"`
	Keywords    string `gorm:"comment:关键字" form:"keywords"`
	Copy        string `gorm:"comment:备案信息" form:"copy"`
	SiteIco     string `gorm:"comment:站点图标" form:"site_ico"`
	SiteImg     string `gorm:"comment:站点图片" form:"site_img"`
	SiteUrl     string `gorm:"comment:站点地址" form:"site_url"`
	Description string `gorm:"comment:描述;" form:"description"`
}

func (o *Option) GetOptionJson() *response.OptionJson {
	return &response.OptionJson{
		Id:          o.Id,
		Title:       o.Title,
		Description: o.Description,
		Keywords:    o.Keywords,
		Copy:        o.Copy,
		SiteIco:     o.SiteIco,
		SiteImg:     o.SiteImg,
		SiteUrl:     o.SiteUrl,
		CreateTime:  utils.FormatTime(o.CreatedAt),
		UpdateTime:  utils.FormatTime(o.UpdatedAt),
	}
}
