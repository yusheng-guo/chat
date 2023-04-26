package service

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/yushengguo557/chat/global"
)

// 注册测试
func TestRegister(t *testing.T) {
	global.InitDB()
	s := NewService(context.TODO())
	err := s.Register(&RegisterRequest{
		// Email:    "declinedyew@outlook.com",
		// Password: "3x+4y=25",
		Email:    "123",
		Password: "123",
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("注册成功！")
	}
}

// 登录测试
func TestLogin(t *testing.T) {
	global.InitDB()
	s := NewService(context.TODO())
	rs := s.Login(&LoginRequest{
		Email:    "declinedyew@outlook.com",
		Password: "3x+4y=25",
	})
	log.Println(rs.Message)
}
