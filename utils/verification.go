package utils

import (
	"github.com/Summrry-top/gin-glory-api/models/request"
	"regexp"
)

const (
	PatternAccount  = `^[a-zA-Z0-9_-]{4,16}$`
	PatternEmail    = `^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$`
	PatternPassword = `^[a-zA-Z0-9]{4,16}$`
	PatternPhone    = `^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0,5-9]))\d{8}$`
	PatternUrl      = `^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$`
	PatternImg      = ``
	PatternName     = ``
	//Pattern
)

// 校验getParam
func VerifyParam(p request.GetParam) bool {
	if p.Id < 0 || p.Page < 1 || p.Limit < 1 || p.Limit > 120 {
		return false
	}
	return true
}

// 校验邮箱
func VerifyEmail(s string) bool {
	b, _ := regexp.MatchString(PatternEmail, s)
	return b
}

// 校验账户
func VerifyAccount(s string) bool {
	b, _ := regexp.MatchString(PatternAccount, s)
	return b
}

// 校验密码
func VerifyPassword(s string) bool {
	b, _ := regexp.MatchString(PatternPassword, s)
	return b
}

func VerifyUrl(s string) bool {
	b, _ := regexp.MatchString(PatternUrl, s)
	return b
}

func VerifyName(s string) bool {
	l := len(s)
	return 0 < l && l <= 50
}

// 为nil
func Nil(n interface{}) bool {
	return n == nil
}

// 非nil
func NotNil(n interface{}) bool {
	return n != nil
}

// 为空
func Empty(s string) bool {
	return s == ""
}

// 非空
func NotEmpty(s string) bool {
	return s != ""
}

// 为零
func Zero(n int) bool {
	return n == 0
}

// 相等
func Equal(a interface{}, b interface{}) bool {
	return a == b
}

// 不相等
func UnEqual(a interface{}, b interface{}) bool {
	return a != b
}
