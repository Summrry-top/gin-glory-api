package initialize

import (
	"fmt"
	"github.com/Summrry-top/gin-glory-api/global"
	"github.com/Summrry-top/gin-glory-api/models/config"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func InitConfig(file string) {
	// 读取配置文件
	v := viper.New()
	v.SetConfigFile(file)
	// 动态监测配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件发生变化！")
	})
	if err := v.ReadInConfig(); err != nil {
		fmt.Println("配置文件致命错误:", err)
		os.Exit(0)
	}
	var server config.Server
	if err := v.Unmarshal(&server); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	global.ServerConfig = server.ServerConfig
	//fmt.Println(file,server,v.Get("server"))
	//fmt.Println(global.ServerConfig)
	global.Viper = v
	// 邮件auth
	//global.SmtpAuth = smtp.PlainAuth("",global.ServerConfig.Smtp.Username,global.ServerConfig.Smtp.Password,global.ServerConfig.Smtp.Host)
	// jwtSecret
	//global.JwtSecret = []byte(global.ServerConfig.Jwt.Key)
}
