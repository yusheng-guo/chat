package utils

import (
	"fmt"
	"testing"
)

// TestGenRandFourDigits 测试随机四位数生成功能
func TestGenRandFourDigits(t *testing.T) {
	foo := GenRandFourDigits()
	for i := 0; i < 10; i++ {
		fmt.Println(foo())
	}
}

// TestSendEmail 测试邮件发送功能
func TestSendEmail(t *testing.T) {
	e := NewEmailService()
	// err := e.Send("declinedyew@outlook.com", "验证码", GenRandFourDigits()())
	err := e.SendCaptcha("declinedyew@outlook.com", e.GenCaptcha())
	if err != nil {
		panic(err)
	}
}
