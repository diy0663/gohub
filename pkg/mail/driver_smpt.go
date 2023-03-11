package mail

import (
	// 基于SMTP协议的第三方邮件包
	"fmt"
	"net/smtp"

	"github.com/diy0663/gohub/pkg/logger"
	emailPKG "github.com/jordan-wright/email"
)

type SMTP struct{}

func (s *SMTP) Send(email Email, config map[string]string) bool {
	e := emailPKG.NewEmail()
	e.From = fmt.Sprintf("%v <%v>", email.From.Name, email.From.Address)
	e.To = email.To
	e.Bcc = email.Bcc
	e.Subject = email.Subject
	e.Text = email.Text
	e.HTML = email.HTML

	logger.DebugJSON("发送邮件", "发送详情", e)
	url := config["host"] + ":" + config["port"]
	err := e.Send(url, smtp.PlainAuth("", config["username"], config["password"], config["host"]))
	if err != nil {
		logger.ErrorString("发送邮件", "发件出错", err.Error())
		return false
	}
	logger.DebugJSON("发送邮件", "发件成功", "")
	return true
}
