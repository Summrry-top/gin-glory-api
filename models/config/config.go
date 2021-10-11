package config

import "time"

type Server struct {
	ServerConfig ServerConfig `mapstructure:"server" yaml:"server"`
}

type ServerConfig struct {
	Install bool  `yaml:"install"`
	App     App   `yaml:"app"`
	Log     Log   `yaml:"log"`
	Mysql   Mysql `yaml:"mysql"`
	Redis   Redis `yaml:"redis"`
	Jwt     Jwt   `yaml:"jwt"`
	Smtp    Smtp  `yaml:"smtp"`
}

// 应用信息
type App struct {
	Mode    string `yaml:"mode"`
	Port    int    `yaml:"port"`
	Author  string `yaml:"author"`
	Git     string `yaml:"git"`
	Version string `yaml:"version"`
}

// 日志信息
type Log struct {
	Path       string     `yaml:"path"`
	Level      string     `yaml:"level"`
	FilePrefix string     `yaml:"filePrefix"`
	FileFormat string     `yaml:"fileFormat"`
	OutFormat  string     `yaml:"outFormat"`
	LumberJack LumberJack `yaml:"lumberJack"`
}

// 日志切割
type LumberJack struct {
	MaxSize    int  `yaml:"maxsize"`
	MaxBackups int  `yaml:"maxBackups"`
	MaxAge     int  `yaml:"maxAge"`
	Compress   bool `yaml:"compress"`
}

// 数据库
type Mysql struct {
	Host         string `yaml:"host" form:"mysql_host"`
	Port         int    `yaml:"port" form:"mysql_port"`
	DbName       string `yaml:"dbName" form:"mysql_name"`
	Username     string `yaml:"username" form:"mysql_username"`
	Password     string `yaml:"password" form:"mysql_password"`
	Config       string `yaml:"config" form:"mysql_config"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
	TablePrefix  string `yaml:"tablePrefix" form:"mysql_table_prefix"` // 表前缀
	//LogMode     bool   `yaml:"logMode"`
	//LogZap      string `yaml:"logZap"`
	//Charset     string `yaml:"charset"`
	//ParseTime   bool   `yaml:"parse_time"`
	//TimeZone    string `yaml:"timeZone"`
}

type Redis struct {
	Addr     string `yaml:"addr" form:"redis_addr"`
	Db       int    `yaml:"db" form:"redis_db"`
	Password string `yaml:"password" form:"redis_password"`
	//PoolSize int    `yaml:"poolSize"`
}

type Jwt struct {
	Key    string        `yaml:"key"`
	Issuer string        `yaml:"issuer"`
	Expire time.Duration `yaml:"expire"`
}
type Smtp struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Nickname string `yaml:"nickname"`
	Password string `yaml:"password"`
}
