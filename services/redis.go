package services

import (
	// "fmt"
	"encoding/json"
	"github.com/Summrry-top/gin-glory-api/global"
	"go.uber.org/zap"
	"strings"
)

//func SetCode(g *request.GetCodeParam) *response.Response {
//	key := GetKey(g.Email, "code")
//	code, err := global.Redis.Get(global.Ctx, key).Result()
//	if err != nil {
//		code = utils.CreateCode(6)
//		global.Redis.Set(global.Ctx, key, code, 15*time.Minute)
//	}
//	// utils.SendSMTPMail(g.Email,"验证码","验证码："+code)
//	return Success("验证码已发送！"+code, global.Nil)
//}

func GetKey(args ...string) string {
	return strings.Join(args, "_")
}

//func

// 记录错误日志
func log(msg, key, field string, err error) {
	global.Logger.Error(msg,
		zap.String("key", key),
		zap.String("field", field),
		zap.String("error", err.Error()),
	)
}

// 设置hash表key中field的value值，记录错误日志
func Hset(key, field string, value interface{}) {
	dest, err := json.Marshal(value)
	if err != nil {
		log("序列化失败！", key, field, err)
	}
	err = global.Redis.HSet(global.Ctx, key, field, dest).Err()
	if err != nil {
		log("设置缓存失败！", key, field, err)
	}

}

// 获取hash表key中field的value并序列化，记录错误日志
func HGet(key, field string, dest interface{}) bool {
	b, err := global.Redis.HGet(global.Ctx, key, field).Bytes()
	if err != nil {
		log("获取缓存失败！", key, field, err)
		return false
	}
	err = json.Unmarshal(b, &dest)
	if err != nil {
		log("反序列化失败！", key, field, err)
		return false
	}
	return true
}

// #string到int
// int,err:=strconv.Atoi(string)
// #string到int64
// int64, err := strconv.ParseInt(string, 10, 64)
// #int到string
// string:=strconv.Itoa(int)
// #int64到string
// string:=strconv.FormatInt(int64,10)
