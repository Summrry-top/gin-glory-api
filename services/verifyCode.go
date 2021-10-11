package services

import (
	"fmt"
	"github.com/Summrry-top/gin-glory-api/global"
	"github.com/Summrry-top/gin-glory-api/models/response"
	"github.com/Summrry-top/gin-glory-api/utils"
	"time"
)

// 生成验证码 有效期15分钟 发送验证码
func CreateVerifyCode(email string) *response.Response {
	if !utils.VerifyEmail(email) {
		return Err400("请输入正确的邮箱")
	}
	// 生成key
	key := GetKey("verify_code", email)
	// 获取redis缓存
	code, err := global.Redis.Get(global.Ctx, key).Result()
	fmt.Println(code, err)
	if err != nil {
		code = utils.CreateCode(6)
		// 设置redis缓存 有效期十五分钟
		global.Redis.Set(global.Ctx, key, code, 15*time.Minute)
	}
	fmt.Println(email)
	// 发送邮件
	msg, ok := utils.SendSMTPMail(email, "验证码", "验证码（有效期15分钟）为："+code)
	if ok {
		return Success(msg, global.Nil)
	}
	return Err400(msg)
}

// 校验验证码
func CheckVerifyCode(email, code string) bool {
	// 生成key
	key := GetKey("verify_code", email)
	// 获取redis缓存
	result, err := global.Redis.Get(global.Ctx, key).Result()
	if err != nil || code != result {
		return false
	}
	return true
}
