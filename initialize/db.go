package initialize

import (
	"fmt"
	"github.com/Summrry-top/gin-glory-api/global"
	"github.com/Summrry-top/gin-glory-api/models/orm"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
)

// mysql
func InitDb() {
	var err error
	// "root:root@tcp(127.0.0.1:3306)/glory?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		global.ServerConfig.Mysql.Username,
		global.ServerConfig.Mysql.Password,
		global.ServerConfig.Mysql.Host,
		global.ServerConfig.Mysql.Port,
		global.ServerConfig.Mysql.DbName,
		global.ServerConfig.Mysql.Config,
	)
	// 连接数据库
	global.Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent), // 日志模式
		SkipDefaultTransaction:                   true,                                  // 禁用默认事务
		DisableForeignKeyConstraintWhenMigrating: true,                                  // 禁用外键约束
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",  // 表名前缀
			SingularTable: false, // 单数表名
		},
	})
	if err != nil {
		global.Logger.Error("连接mysql数据库失败",
			zap.String("error", err.Error()),
		)
		fmt.Println("连接数据库失败", err.Error())
		os.Exit(0)
	}
	// 同步表
	err = global.Db.AutoMigrate(new(orm.User),
		new(orm.Link), new(orm.LinkSort),
		new(orm.Banner), new(orm.Music), new(orm.Option),
		new(orm.Article), new(orm.ArticleSort), new(orm.ArticleTag))
	if err != nil {
		global.Logger.Error("表同步失败",
			zap.String("error", err.Error()),
		)
		fmt.Println(err)
	}
}

// 连接mysql
func ConnectDb() bool {
	m := global.ServerConfig.Mysql
	// "root:root@tcp(127.0.0.1:3306)/glory?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.DbName,
		m.Config,
	)
	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent), // 日志级别
		SkipDefaultTransaction:                   true,                                  // 禁用默认事务
		DisableForeignKeyConstraintWhenMigrating: true,                                  // 禁用外键约束
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   m.TablePrefix, // 表名前缀
			SingularTable: false,         // 单数表名
		},
	})
	if err != nil {
		fmt.Println("连接数据库失败", err.Error())
		return false
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	global.Db = db
	return true
}

// 同步表
func MigrateDb() bool {
	// 同步表
	err := global.Db.AutoMigrate(new(orm.User),
		new(orm.Link), new(orm.LinkSort),
		new(orm.Banner), new(orm.Music), new(orm.Option),
		new(orm.Article), new(orm.ArticleSort), new(orm.ArticleTag))
	if err != nil {
		global.Logger.Error("表同步失败",
			zap.String("error", err.Error()),
		)
		fmt.Println(err)
		return false
	}
	return true
}
