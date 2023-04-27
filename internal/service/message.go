package service

import (
	"fmt"
	"net"

	"github.com/yushengguo557/chat/internal/model"
)

type SendMessageRequest struct {
	Sender   string            `form:"sender"`   // 发送者
	Receiver string            `form:"receiver"` // 接收者
	Content  string            `form:"content"`  // 消息内容
	Type     model.MessageType `form:"type"`     // 消息类型
}

// SendMessage 用户发送消息
func (s *Service) SendMessage(smr *SendMessageRequest) error {
	// 1.接收者是否在线
	if s.dao.IsOnline(smr.Receiver) {
		receiver, err := s.dao.GetOnlineUser(smr.Receiver)
		if err != nil {
			return fmt.Errorf("sending message to user: %w", err)
		}
		msg := model.NewMessage(smr.Sender, smr.Receiver, smr.Content, smr.Type)
		var serialized []byte
		if serialized, err = msg.Marshal(); err != nil {
			return fmt.Errorf("marshaling message in `SendMessage Function`: %w", err)
		}
		// // 2.用户在线 发送消息
		receiver.Conn.Write([]byte(serialized))

		// 3.将消息存储到rethinkdb数据库
		err = s.dao.InsertMessage(msg)
		if err != nil {
			return fmt.Errorf("inserting message to rethinkdb in `SendMessage Function`: %w", err)
		}
		return nil
	}
	// 4.用户不在线 将消息存放在redis哈希表中
	// 5.遍历redis中的消息
	//	- 若用户上线了 则发送消息
	//	- 若消息自发出后 十分钟 接收者未上线 则将消息存储到rethinkdb 发送但未接收消息表中
	return nil
}

// ReceiveMessage 用户接收消息
func (s *Service) ReceiveMessage() *model.Message {
	// 从 rethinkdb 中返回所有用户未读取的消息

	return nil
}

// ReceiveAndSend接收并发送消息
func (s *Service) ReceiveAndSend(conn net.Conn) {
}
