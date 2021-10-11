package services

import (
	"github.com/Summrry-top/gin-glory-api/models/orm"
	"github.com/Summrry-top/gin-glory-api/models/request"
	"github.com/Summrry-top/gin-glory-api/models/response"
	"github.com/Summrry-top/gin-glory-api/utils"
	"strconv"
)

// 获取单个user 通过id
func GetUserById(isAdmin bool, id int) *response.Response {
	var user orm.User
	if ok := GetOne("id", id, &user); !ok {
		return Err400("用户不存在！")
	}
	if isAdmin {
		return Success("管理员单条数据", user.GetUserAdminJson())
	}
	return Success("单条数据", user.GetUserJson())
}

// 获取全部user
func GetUserAll(isAdmin bool, p request.GetParam) *response.Response {
	var users []orm.User
	var pageData response.PageData
	if !GetPagination(p, &pageData, &users) {
		return Err400("没有数据！")
	}
	if isAdmin {
		var usersAdminJson []orm.UserAdminJson
		for _, v := range users {
			usersAdminJson = append(usersAdminJson, v.GetUserAdminJson())
		}
		pageData.Data = usersAdminJson
		return Success("管理员多条数据", pageData)
	}
	var usersJson []orm.UserJson
	for _, v := range users {
		usersJson = append(usersJson, v.GetUserJson())
	}
	pageData.Data = usersJson
	return Success("多条数据", pageData)
}

// 登录
func Login(p *orm.User) *response.Response {
	if !utils.VerifyEmail(p.Email) {
		return Err400("邮箱不合法")
	}
	if !utils.VerifyPassword(p.Password) {
		return Err400("密码不合法")
	}
	var u orm.User
	if !GetOne("email", p.Email, &u) {
		return Err400("邮箱不存在")
	}
	if !utils.CheckPw(u.Password, p.Password) {
		return Err400("密码错误！")
	}
	var loginData response.LoginData
	isAdmin := u.Level == "admin"
	loginData.LoginToken, _ = utils.CreateToken(u.Account, isAdmin)
	if isAdmin {
		loginData.User = u.GetUserAdminJson()
		return Success("管理员登录成功", loginData)
	}
	loginData.User = u.GetUserJson()
	return Success("登录成功！", loginData)
}

// 注册
func Register(p *orm.User, code string) *response.Response {
	if !utils.VerifyEmail(p.Email) {
		return Err400("邮箱不合法！")
	}
	if !CheckVerifyCode(p.Email, code) {
		return Err400("验证码错误")
	}
	var user orm.User
	if GetOne("email", p.Email, &user) {
		return Err400("邮箱已存在！")
	}
	if GetOne("account", p.Account, &user) {
		return Err400("账户已存在！")
	}
	if CreateOne(p) {
		var loginData response.LoginData
		loginData.LoginToken, _ = utils.CreateToken(user.Account, false)
		loginData.User = user.GetUserJson()
		return Success("注册成功", loginData)
	}
	return Err400("注册失败")
}

// 修改
func PostUserEdit(p *orm.User) *response.Response {
	if utils.Zero(p.Id) {
		return Err400("id错误")
	}
	if UpdateOne(p) {
		return Ok()
	}
	return Err400("文章更新失败")
}

// 删除
func PostUserDelete(idStr string) *response.Response {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return Err400("id错误")
	}
	if DeleteOne(id, new(orm.Article)) {
		return Ok()
	}
	return Err400("文章删除失败")
}
