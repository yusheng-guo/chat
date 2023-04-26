package service

import "github.com/yushengguo557/chat/internal/model"

type SendMessageRequest struct {
	Sender   string            `form:"sender"`   // 发送者
	Receiver string            `form:"receiver"` // 接收者
	Content  string            `form:"content"`  // 消息内容
	Type     model.MessageType `form:"type"`     // 消息类型
}

// SendMessage 用户发送消息
func (s *Service) SendMessage(mr *SendMessageRequest) {
	// 1.接收者是否在线
	// 2.用户在线 发送消息
	// 3.将消息存储到rethinkdb数据库

	// 4.用户不在线 将消息存放在redis哈希表中
	// 5.遍历redis中的消息
	//	- 若用户上线了 则发送消息
	//	- 若消息自发出后 十分钟 接收者未上线 则将消息存储到rethinkdb 发送但未接收消息表中
}

// ReceiveMessage 用户接收消息
func (s *Service) ReceiveMessage() []*model.Message {
	// 从 rethinkdb 中返回所有用户未读取的消息

	return nil
}
