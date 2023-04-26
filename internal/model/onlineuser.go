package model

import (
	"github.com/yushengguo557/chat/global"
)

// OnlineUser 在线用户
type OnlineUser struct {
	ID string `json:"id"` // 用户ID
	// IO   sync.Mutex         `json:"io"`   // 互斥锁
	// Conn io.ReadWriteCloser `json:"conn"` // 用户读写连接
}

// NewOnlineUser 实例化 OnlineUser 结构体
func NewOnlineUser(id string) *OnlineUser {
	return &OnlineUser{
		ID: id,
	}
}

// 序列化在线用户
func (o *OnlineUser) Marshal() ([]byte, error) {
	return global.Json.Marshal(o)
}

// 反序列化在线用户
func (o *OnlineUser) Unmarshal(data []byte) error {
	return global.Json.Unmarshal(data, &o)
}
