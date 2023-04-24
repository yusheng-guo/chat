package service

import (
	"context"
	"testing"
)

func TestRegister(t *testing.T) {
	// s := New(nil)
	s := New(context.TODO())
	s.Register(&RegisterRequest{
		Email:    "yushengguo557@gmail.com",
		Password: "3x+4y=25",
	})
}
