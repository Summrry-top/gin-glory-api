package initialize

import (
	"fmt"
	"github.com/Summrry-top/gin-glory-api/global"
	"net/smtp"
)

func InitSmtp() {
	global.SmtpAddr = fmt.Sprintf("%s:%d", global.ServerConfig.Smtp.Host, global.ServerConfig.Smtp.Port)
	// 邮件auth
	global.SmtpAuth = smtp.PlainAuth("", global.ServerConfig.Smtp.Username, global.ServerConfig.Smtp.Password, global.ServerConfig.Smtp.Host)
}
