package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/yushengguo557/chat/global"
)

func TestRegister(t *testing.T) {
	global.InitDB()
	s := NewService(context.TODO())
	err := s.Register(&RegisterRequest{
		Email:    "declinedyew@outlook.com",
		Password: "3x+4y=25",
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("注册成功！")
	}
}
