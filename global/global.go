package global

import (
	"context"
	"github.com/Summrry-top/gin-glory-api/models/config"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/smtp"
)

const (
	// 配置文件
	ConfigFile        = "./config/config.yaml"
	DefaultConfigFile = "./config/default.yaml"
	// cors
	AllowCredentials = "true"
	AllowMethods     = "POST, GET"
	AllowHeaders     = "Origin, X-Requested-With, Content-Type, Accept, Authorization"
	ExposeHeaders    = "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type"
	// smtp
	SmtpContentType = "Content-Type: text/html; charset=UTF-8"
	// 请求参数默认值
	Limit = 5
	Order = "created_at"
	Page  = 1
	Cache = true
	// redis
	RedisUser   = "hash_user"
	VerticalBar = "|"
	Hyphen      = "-"
	Underline   = "_"
	LikeLeft    = "%|"
	LikeRight   = "|%"

	//
	OK = 200
)

var (
	// 配置
	ServerConfig config.ServerConfig
	Viper        *viper.Viper
	// 日志
	Logger *zap.Logger
	// mysql
	Db *gorm.DB
	// redis
	Redis *redis.Client
	// smtp
	SmtpAuth smtp.Auth
	SmtpAddr string
	// jwt
	JwtSecret []byte

	// 空数组
	Nil = make([]int, 0)
	// 维护分类标签id
	Array = []string{LikeLeft, LikeRight}
	// 空上下文
	Ctx = context.Background() // go-redis必须
	// 初始化存储器
	Store = sessions.NewCookieStore([]byte("SESSION_KEY"))
)
