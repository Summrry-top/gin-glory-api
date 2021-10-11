package initialize

import (
	"github.com/Summrry-top/gin-glory-api/global"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path"
	"strings"
	"time"
)

// 初始化日志
func InitLogger() {
	// 日志配置变量
	var logConfig = global.ServerConfig.Log
	var encoder zapcore.Encoder
	if logConfig.OutFormat == "json" {
		encoder = zapcore.NewJSONEncoder(getEncoderConfig())
	} else {
		encoder = zapcore.NewConsoleEncoder(getEncoderConfig())
	}
	// 设置日志文件切割和归档
	WriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   getLogFile(),                    // 日志文件
		MaxSize:    logConfig.LumberJack.MaxAge,     // 单文件最大容量(MB)
		MaxBackups: logConfig.LumberJack.MaxBackups, // 保留旧文件的最大数量
		MaxAge:     logConfig.LumberJack.MaxAge,     //  旧文件最多保存几天
		Compress:   logConfig.LumberJack.Compress,   // 是否压缩/归档旧文件
	})
	// 创建NewCore
	zapCore := zapcore.NewCore(encoder, WriteSyncer, getLevel())
	// 创建logger
	logger := zap.New(zapCore)
	//defer logger.Sync()
	// 赋值给全局变量
	global.Logger = logger
}

// 获取最低记录日志级别
func getLevel() zapcore.Level {
	// 日志级别map
	var levelMap = map[string]zapcore.Level{
		"debug":  zapcore.DebugLevel,
		"info":   zapcore.InfoLevel,
		"warn":   zapcore.WarnLevel,
		"error":  zapcore.ErrorLevel,
		"dPanic": zapcore.DPanicLevel,
		"panic":  zapcore.PanicLevel,
		"fatal":  zapcore.FatalLevel,
	}
	if level, ok := levelMap[global.ServerConfig.Log.Level]; ok {
		return level
	}
	return zapcore.InfoLevel
}

// 自定义日志输出格式
func getEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		FunctionKey:   zapcore.OmitKey,
		MessageKey:    "msg",
		StacktraceKey: "s",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime:    getEncodeTime,
	}

}

// 定义日志输出时间格式
func getEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 - 15:04:05.000"))
}

// // 获取文件切割和归档配置信息
// func getLumberjackWriteSyncer() {

// }

// 获取日志文件名
func getLogFile() string {
	fileFormat := time.Now().Format(global.ServerConfig.Log.FileFormat)
	fileName := strings.Join([]string{
		global.ServerConfig.Log.FilePrefix,
		fileFormat,
		"log",
	}, ".")
	return path.Join(global.ServerConfig.Log.Path, fileName)
}
