package router

import (
	"fmt"
	"github.com/Summrry-top/gin-glory-api/api"
	"github.com/Summrry-top/gin-glory-api/global"
	"github.com/Summrry-top/gin-glory-api/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
)

func InitRouter() {
	// mode
	gin.SetMode(global.ServerConfig.App.Mode)
	r := gin.New()
	// 请求日志写入文件
	r.Use(middleware.GinLogger())
	// 统一错误处理
	r.Use(middleware.Recovery)
	// 跨域
	r.Use(middleware.Cors())
	// 404
	r.NoRoute(api.Err404)
	// 默认路由组
	DefaultRouteGroup(r)
	// 运行
	addr := fmt.Sprintf(":%d", global.ServerConfig.App.Port)
	err := r.Run(addr)
	if err != nil {
		global.Logger.Error("路由器启动失败",
			zap.String("err", err.Error()))
		os.Exit(0)
	}

}

func DefaultRouteGroup(r *gin.Engine) {
	// api
	d := r.Group("api")
	d.POST("/install", api.PostInstall)
	d.Use(middleware.Install())
	{
		d.GET("/article", api.GetArticle)
		d.GET("/articleSort", api.GetArticleSort)
		d.GET("/tag", api.GetArticleTag)
		d.GET("/link", api.GetLink)
		d.GET("/linkSort", api.GetLinkSort)
		d.GET("/banner", api.GetBanner)
		d.GET("/option", api.GetOption)
		d.POST("/verifyCode", api.PostVerifyCode)
	}
	// jwt
	d.Use(middleware.MyJwtToken())
	{
		d.GET("/user", api.GetUser)
		d.POST("/user", api.PostUser)
	}
	// 权限校验
	d.Use(middleware.Auth())
	{
		d.POST("/article", api.PostArticle)
		d.POST("/articleSort", api.PostArticleSort)
		d.POST("/tag", api.PostArticleTag)
		d.POST("/link", api.PostLink)
		d.POST("/linkSort", api.PostLinkSort)
		d.POST("/banner", api.PostBanner)
		d.POST("/option", api.PostOption)
	}

}
