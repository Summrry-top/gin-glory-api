package services

import (
	"fmt"
	"github.com/Summrry-top/gin-glory-api/global"
	"github.com/Summrry-top/gin-glory-api/initialize"
	"github.com/Summrry-top/gin-glory-api/models/orm"
	"github.com/Summrry-top/gin-glory-api/utils"

	//"github.com/Summrry-top/gin-glory-api/initialize"
	"github.com/Summrry-top/gin-glory-api/models/config"
	"github.com/Summrry-top/gin-glory-api/models/response"
	"github.com/gin-gonic/gin"
)

// 测试数据库连接
func CheckDb(c *gin.Context) *response.Response {
	//var dbInfo config.Mysql
	//_=c.ShouldBind(&dbInfo)
	//if !CheckDbInfo(dbInfo){
	//	Err400("数据库信息不合法")
	//}
	//global.ServerConfig.Mysql=dbInfo
	_ = c.ShouldBind(&global.ServerConfig.Mysql)
	fmt.Println(global.ServerConfig.Mysql)
	//global.ServerConfig.Mysql.Host=c.PostForm("host")
	//global.ServerConfig.Mysql.Port=c.PostForm("port")
	//global.ServerConfig.Mysql.Username=c.PostForm("username")
	//global.ServerConfig.Mysql.Password=c.PostForm("password")
	//global.ServerConfig.Mysql.DbName=c.PostForm("db_name")
	//global.ServerConfig.Mysql.Config=c.PostForm("config")
	if initialize.ConnectDb() {
		return Success("数据库连接成功", global.Nil)
	}
	return Err400("连接数据库失败")
}

// 测试redis连接
func CheckRedis(c *gin.Context) *response.Response {
	//var redisInfo config.Redis
	//_=c.ShouldBind(redisInfo)
	//global.ServerConfig.Redis=redisInfo
	//c.Set("redisInfo",redisInfo)
	//if !CheckRedisInfo(redisInfo){
	//	Err400("redis信息不合法")
	//}
	global.ServerConfig.Redis.Addr = c.PostForm("addr")
	global.ServerConfig.Redis.Db = 0
	if initialize.ConnectRedis() {
		return Success("redis连接成功！", global.Nil)
	}
	return Err400("连接redis失败")

}

// 进行安装
func Install(c *gin.Context) *response.Response {
	data := CheckDb(c)
	if data.Code == 400 {
		return data
	}
	data = CheckRedis(c)
	if data.Code == 400 {
		return data
	}
	if !initialize.MigrateDb() {
		return Err400("生成表失败！")
	}
	var user orm.User
	//_=c.ShouldBind(&user)
	//user.Password=utils.CreatePw(user.Password)
	user.Level = "admin"
	user.Account = "admin"
	user.Password = utils.CreatePw("123456")
	if !CreateOne(&user) {
		return Err400("初始化管理员信息失败")
	}
	//写入配置信息
	fmt.Println(global.ServerConfig, 1)
	global.ServerConfig.Install = true
	global.Viper.Set("server", global.ServerConfig)
	err := global.Viper.WriteConfigAs(global.ConfigFile)
	if err != nil {
		global.ServerConfig.Install = false
		return Err400("配置信息写入失败")
	}
	fmt.Println(global.ServerConfig)
	return Success("安装成功！", global.Nil)
}

func CheckDbInfo(dbInfo config.Mysql) bool {
	return true
}

func CheckRedisInfo(redisInfo config.Redis) bool {
	return true
}
