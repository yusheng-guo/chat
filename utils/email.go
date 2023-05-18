package utils

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/wneessen/go-mail"
	"github.com/yushengguo557/chat/global"
)

type EmailService struct {
	*mail.Client
	GenCaptcha func() string
}

type Data struct {
	Code string
}

// NewEmailService 实例化邮件服务
func NewEmailService() (*EmailService, error) {
	// cli, err := mail.NewClient("smtp-mail.outlook.com",
	// 	mail.WithPort(587), mail.WithSMTPAuth(mail.SMTPAuthLogin),
	// 	mail.WithUsername("declinedyew@outlook.com"), mail.WithPassword("3x+4y=25"))
	host := global.Email.Host
	port := global.Email.Port
	username := global.Email.Username
	password := global.Email.Password
	cli, err := mail.NewClient(host,
		mail.WithPort(port), mail.WithSMTPAuth(mail.SMTPAuthLogin),
		mail.WithUsername(username), mail.WithPassword(password))
	if err != nil {
		return nil, err
	}

	return &EmailService{
		Client:     cli,
		GenCaptcha: GenRandFourDigits(),
	}, nil
}

// Close 关闭服务
// func (e *EmailService) Close() error {
// 	return e.Client.Close().Error()
// }

// Send 发送邮件
func (e *EmailService) Send(from, to, subject, content string) (err error) {
	m := mail.NewMsg()
	if err = m.From(from); err != nil {
		return err
	}
	if err = m.To(to); err != nil {
		return err
	}
	m.Subject(subject)
	m.SetBodyString(mail.TypeTextHTML, content)
	if err = e.Client.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

// SendCaptcha 发送四位随机数验证码邮件
func (e *EmailService) SendCaptcha(to string) error {
	// 解析 html文件
	t, err := template.ParseFiles("captcha.html")
	// t, err := template.ParseFiles("assets\\html\\captcha\\captcha.html")
	if err != nil {
		return fmt.Errorf("new and parse template: %w", err)
	}

	// 渲染数据
	data := Data{
		Code: e.GenCaptcha(),
	}
	var buff bytes.Buffer
	if err = t.Execute(&buff, data); err != nil {
		return fmt.Errorf("execute data: %w", err)
	}

	// 发送邮件
	if e.Send("declinedyew@outlook.com", to, "验证码", buff.String()); err != nil {
		return fmt.Errorf("send email: %w", err)
	}
	return nil
}

// 基于 "gopkg.in/gomail.v2" 库的实现 但该库已长时间未更新
// type EmailService struct {
// 	D          *gomail.Dialer
// 	GenCaptcha func() string
// }

// // InitEmail 初始化 Email 服务
// func NewEmailService() *EmailService {
// 	return &EmailService{
// 		D:          gomail.NewDialer("smtp-mail.outlook.com", 587, "declinedyew@outlook.com", "3x+4y=25"),
// 		GenCaptcha: GenRandFourDigits(),
// 	}
// }

// // Send 发送邮件
// func (e *EmailService) Send(to, subject, content string) error {
// 	m := gomail.NewMessage()
// 	m.SetHeader("From", "declinedyew@outlook.com")
// 	m.SetHeader("To", to)
// 	m.SetHeader("Subject", subject)
// 	m.SetBody("text/html", fmt.Sprintf("<p>%s</p>", content))

// 	if err := e.D.DialAndSend(m); err != nil {
// 		return fmt.Errorf(": %w", err)
// 	}
// 	return nil
// }

// // SendCaptcha 发送四位随机数验证码邮件
// func (e *EmailService) SendCaptcha(to, content string) error {
// 	return e.Send(to, "验证码", e.GenCaptcha())
// }
