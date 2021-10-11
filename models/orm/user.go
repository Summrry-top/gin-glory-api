package orm

import (
	"github.com/Summrry-top/gin-glory-api/utils"
	"gorm.io/datatypes"
)

// orm模型
type User struct {
	Model
	Account     string         `gorm:"comment:账户;unique" form:"account"`
	Password    string         `gorm:"comment:密码;" form:"password"`
	Nickname    string         `gorm:"comment:昵称;" form:"nickname"`
	Email       string         `gorm:"comment:邮箱;unique_index;" form:"email"`
	Phone       string         `gorm:"comment:电话;" form:"phone"`
	Sex         string         `gorm:"comment:性别;default:保密" form:"sex"`
	Status      int8           `gorm:"comment:状态;default:1;" form:"status"`
	Level       string         `gorm:"comment:等级;default:user;" form:"level"`
	Remarks     string         `gorm:"comment:标记" form:"remarks"`
	HeadImg     string         `gorm:"comment:头像" form:"head_img"`
	Description string         `gorm:"comment:描述" form:"description"`
	AddressUrl  string         `gorm:"comment:地址" form:"address_url"`
	Opt         datatypes.JSON `gorm:"comment:选项" form:"opt"`
}

// 非管理员用户返回的json
type UserJson struct {
	Id            int            `json:"id"`
	Nickname      string         `json:"nickname"`
	Sex           string         `json:"sex"`
	HeadImg       string         `json:"head_img"`
	Description   string         `json:"description"`
	Status        int8           `json:"status"`
	AddressUrl    string         `json:"address_url"`
	Opt           datatypes.JSON `json:"opt"`
	CreateTime    string         `json:"create_time"`
	UpdateTime    string         `json:"update_time"`
	LastLoginTime int            `json:"last_login_time"`
}

// UserAdminJson 管理员用户返回的json
type UserAdminJson struct {
	Id            int            `json:"id"`
	Account       string         `json:"account"`
	Password      string         `json:"password"`
	Nickname      string         `json:"nickname"`
	Email         string         `json:"email"`
	Phone         string         `json:"phone"`
	Sex           string         `json:"sex"`
	Status        int8           `json:"status"`
	Level         string         `json:"level"`
	Remarks       string         `json:"remarks"`
	HeadImg       string         `json:"head_img"`
	Description   string         `json:"description"`
	AddressUrl    string         `json:"address_url"`
	Opt           datatypes.JSON `json:"opt"`
	CreateTime    string         `json:"create_time"`
	UpdateTime    string         `json:"update_time"`
	LastLoginTime int64          `json:"last_login_time"`
}

// 普通用户json
func (u *User) GetUserJson() UserJson {
	return UserJson{
		Id:          u.Id,
		Nickname:    u.Nickname,
		Sex:         u.Sex,
		HeadImg:     u.HeadImg,
		Description: u.Description,
		Status:      u.Status,
		AddressUrl:  u.AddressUrl,
		Opt:         u.Opt,
		CreateTime:  utils.FormatTime(u.CreatedAt),
		UpdateTime:  utils.FormatTime(u.UpdatedAt),
	}
}

// 管理员json
func (u *User) GetUserAdminJson() UserAdminJson {
	return UserAdminJson{
		Id:          u.Id,
		Account:     u.Account,
		Password:    u.Password,
		Nickname:    u.Nickname,
		Sex:         u.Sex,
		HeadImg:     u.HeadImg,
		Description: u.Description,
		Status:      u.Status,
		Level:       u.Level,
		Remarks:     u.Remarks,
		AddressUrl:  u.AddressUrl,
		Opt:         u.Opt,
		CreateTime:  utils.FormatTime(u.CreatedAt),
		UpdateTime:  utils.FormatTime(u.UpdatedAt),
	}
}

// 用户json 普通用户/管理员
func (u *User) GetUser() interface{} {
	if u.Level == "admin" {
		return u.GetUserAdminJson()
	}
	return u.GetUserJson()
}
