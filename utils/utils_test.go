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
	// 新建 Email服务
	e, err := NewEmailService()
	if err != nil {
		panic(fmt.Errorf("new...%w", err))
	}

	// 关闭 客户端
	defer func() {
		fmt.Println("close...")
		_ = e.Client.Close()
	}()

	// err := e.Send("declinedyew@outlook.com", "验证码", GenRandFourDigits()())
	// 发送 验证码
	err = e.SendCaptcha("declinedyew@outlook.com")
	if err != nil {
		panic(fmt.Errorf("send... %w", err))
	}
}
