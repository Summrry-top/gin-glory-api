package utils

import (
	"fmt"
	"github.com/Summrry-top/gin-glory-api/global"
	"net/smtp"
	"strings"
)

func SendSMTPMail(mailAddress string, subject string, body string) (string, bool) {
	msg := []byte("To: " + mailAddress + "\r\nFrom: " +
		global.ServerConfig.Smtp.Nickname + "<" + global.ServerConfig.Smtp.Username + ">\r\nSubject: " +
		subject + "\r\n" + global.SmtpContentType + "\r\n\r\n" + body)

	sendTo := strings.Split(mailAddress, ",")
	if err := smtp.SendMail(
		global.SmtpAddr,
		global.SmtpAuth,
		global.ServerConfig.Smtp.Username,
		sendTo,
		msg); err != nil {

		fmt.Println(mailAddress, err)
		return "邮件发送失败！", false
	}
	return "邮件发送成功！", true

}
