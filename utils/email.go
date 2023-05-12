package utils

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

type EmailService struct {
	D          *gomail.Dialer
	GenCaptcha func() string
}

// InitEmail 初始化 Email 服务
func NewEmailService() *EmailService {
	return &EmailService{
		D:          gomail.NewDialer("smtp-mail.outlook.com", 587, "declinedyew@outlook.com", "3x+4y=25"),
		GenCaptcha: GenRandFourDigits(),
	}
}

// Send 发送邮件
func (e *EmailService) Send(to, subject, content string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "declinedyew@outlook.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", fmt.Sprintf("<p>%s</p>", content))

	if err := e.D.DialAndSend(m); err != nil {
		return fmt.Errorf(": %w", err)
	}
	return nil
}

// SendCaptcha 发送四位随机数验证码邮件
func (e *EmailService) SendCaptcha(to, content string) error {
	return e.Send(to, "验证码", e.GenCaptcha())
}
