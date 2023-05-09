package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/yushengguo557/chat/global"
	"github.com/yushengguo557/chat/internal/model"
)

// type MessageRequestResponse struct {
// 	Sender   string            `form:"sender"`   // 发送者
// 	Receiver string            `form:"receiver"` // 接收者
// 	Content  string            `form:"content"`  // 消息内容
// 	Type     model.MessageType `form:"type"`     // 消息类型
// }

// type ErrResponse struct {
// 	Code    int    `form:"code"`    // 状态码
// 	Message string `form:"message"` // 响应消息
// }

// // SendMessage 将收到的的消息发送到目标客户端
// func (s *Service) SendMessage(msr *SendMessageRequest) error {
// 	// 1.接收者是否在线
// 	if s.dao.IsOnline(msr.Receiver) {
// 		receiver, err := s.dao.GetOnlineUser(msr.Receiver)
// 		if err != nil {
// 			return fmt.Errorf("sending message to user: %w", err)
// 		}
// 		msg := model.NewMessage(msr.Sender, msr.Receiver, msr.Content, msr.Type)
// 		var serialized []byte
// 		if serialized, err = msg.Marshal(); err != nil {
// 			return fmt.Errorf("marshaling message in `SendMessage Function`: %w", err)
// 		}
// 		// // 2.用户在线 发送消息
// 		receiver.Conn.Write([]byte(serialized))

// 		// 3.将消息存储到rethinkdb数据库
// 		err = s.dao.InsertMessage(msg)
// 		if err != nil {
// 			return fmt.Errorf("inserting message to rethinkdb in `SendMessage Function`: %w", err)
// 		}
// 		return nil
// 	}
// 	// 4.用户不在线 将消息存放在redis哈希表中
// 	// 5.遍历redis中的消息
// 	//	- 若用户上线了 则发送消息
// 	//	- 若消息自发出后 十分钟 接收者未上线 则将消息存储到rethinkdb 发送但未接收消息表中
// 	return nil
// }

// ReceiveMessage 接收客户端发送到服务器的消息
// func (s *Service) ReceiveMessage(conn net.Conn) *model.Message {
// 	// 从 rethinkdb 中返回所有用户未读取的消息
// 	return nil
// }

// ReceiveAndSend 从sourceConn接收消息 并 向targetConn发送消息
func (s *Service) ReceiveAndSend(sourceConn, targetConn net.Conn) error {
	var (
		state  = ws.StateServerSide
		reader = wsutil.NewReader(sourceConn, state)
		writer = wsutil.NewWriter(targetConn, state, ws.OpText)
	)
	for {
		header, err := reader.NextFrame()
		if err != nil {
			// handle error
			return err
		}

		// Reset writer to write frame with right operation code.
		writer.Reset(sourceConn, state, header.OpCode)

		if _, err = io.Copy(writer, reader); err != nil {
			// handle error
			return err
		}
		if err = writer.Flush(); err != nil {
			// handle error
			return err
		}
	}
}

func (s *Service) Communicate(conn net.Conn) error {
	// 1.reader 和 decoder
	r := wsutil.NewReader(conn, ws.StateServerSide)
	decoder := json.NewDecoder(r)

	for {
		// 2.准备从 r中读取下一条消息
		hdr, err := r.NextFrame()
		if hdr.OpCode == ws.OpClose {
			fmt.Println("关闭连接")
			return nil
		}
		if err != nil {
			return fmt.Errorf("get next frame: %w", err)
		}

		// 3.解码消息
		// var msg model.Message                 // 用于 存储 的消息
		var transmsg model.TransferredMessage // 用户 传输 的消息
		// if err = decoder.Decode(&msg); err != nil {
		// 	return fmt.Errorf("decode message in `Communicate Function`: %w", err)
		// }
		// fmt.Println(msg)
		if err = decoder.Decode(&transmsg); err != nil {
			return fmt.Errorf("decode transferred message in `Communicate Function`: %w", err)
		}
		fmt.Println("transmsg", transmsg)

		// 4.将消息存储到rethinkdb数据库
		// err = s.dao.InsertMessage(&msg)
		// if err != nil {
		// 	return fmt.Errorf("inserting message to rethinkdb in `Communicate Function`: %w", err)
		// }
		// fmt.Println("消息保存成功", msg)
		// fmt.Println("sender: ", msg.Sender)
		// fmt.Println("reciever: ", msg.Receiver)

		// 5.从redis中获取接收者
		// receiver, err := s.dao.GetOnlineUser(msg.Receiver)
		// if err != nil {
		// 	return fmt.Errorf("receiver not exists")
		// }
		// 从全局变量OnlineUser中获取接收者
		global.Lock.RLock() // 加锁
		receiver, ok := global.OnlineUsers[transmsg.Receiver]
		global.Lock.RUnlock()
		if !ok {
			// return fmt.Errorf("receiver not exists")
			fmt.Println("receiver not exists")
		}

		// 6.消息转发
		w := wsutil.NewWriter(receiver, ws.StateServerSide, ws.OpText)
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(&transmsg); err != nil {
			return fmt.Errorf("encode message in `Communicate Function`: %w", err)
		}
		if err = w.Flush(); err != nil {
			return fmt.Errorf("flush to writer in `Communicate Function`: %w", err)
		}
	}
}
