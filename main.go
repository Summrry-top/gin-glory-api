package main

import (
	"fmt"
	"github.com/Summrry-top/gin-glory-api/global"
	"github.com/Summrry-top/gin-glory-api/initialize"
	"github.com/Summrry-top/gin-glory-api/router"
)

// 1、自定义的校验方法
// db.Order(time asc) 升序
// db.Order(age desc) 降序
func init() {
}
func main() {
	// 读取配置信息
	initialize.InitConfig(global.ConfigFile)
	// 是否已安装
	if global.ServerConfig.Install {
		// 连接mysql
		initialize.ConnectDb()
		// 连接redis
		initialize.ConnectRedis()
		// 发送邮件auth
		//initialize.InitSmtpAuth()
	} else {
		initialize.InitConfig(global.DefaultConfigFile)
	}
	// logger
	initialize.InitLogger()
	fmt.Println(global.ServerConfig)
	// 启动路由器
	router.InitRouter()
}
