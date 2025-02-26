package mail

import (
	"github.com/wuvikr/gohub01/pkg/config"
	"github.com/wuvikr/gohub01/pkg/logger"
	"gopkg.in/gomail.v2"
)

type SMTP struct {
}

func (s *SMTP) Send(email Email) error {
	// 1. 获取 SMTP 配置
	host := config.MustGet[string]("mail.smtp.host")
	port := config.MustGet[int]("mail.smtp.port")
	username := config.MustGet[string]("mail.smtp.username")
	password := config.MustGet[string]("mail.smtp.password")

	// 2. 创建邮件消息
	m := gomail.NewMessage()
	m.SetHeader("From", email.From.Address)
	m.SetHeader("To", email.To...)
	if len(email.Cc) > 0 {
		m.SetHeader("Cc", email.Cc...)
	}
	if len(email.Bcc) > 0 {
		m.SetHeader("Bcc", email.Bcc...)
	}
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/html", string(email.HTML))
	// m.Attach("/home/Alex/lolcat.jpg")

	logger.DebugJson("发送邮件信息", "Message", m)

	// 3. 连接到 SMTP 服务器发送邮件
	d := gomail.NewDialer(host, port, username, password)
	if err := d.DialAndSend(m); err != nil {
		logger.ErrorStr("发送邮件失败", err.Error())
		return err
	}

	logger.Info("发送邮件成功")
	return nil
}
