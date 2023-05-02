package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

// MessageStatus 消息状态
type MessageStatus int

const (
	MessageStatusFailed   MessageStatus = iota // 消息发送失败
	MessageStatusSent                          // 已发送
	MessageStatusReceived                      // 已接收
	MessageStatusRead                          // 已读取
)

// MessageType 消息类型
type MessageType int

const (
	MessageTypeText  MessageType = iota // 文本
	MessageTypeImage                    // 图片
	MessageTypeVideo                    // 视频
	MessageTypeFile                     // 文件
	MessageTypeAudio                    // 音频
	MessageTypeMD                       // markdown
)

type Model struct {
	ID        string     `json:"id" gorethink:"id"`                           // 模型ID
	CreatedAt *time.Time `json:"created_at" gorethink:"created_at,omitempty"` // 创建时间
	UpdatedAt *time.Time `json:"updated_at" gorethink:"updated_at,omitempty"` // 更新时间
	DeletedAt *time.Time `json:"deleted_at" gorethink:"deleted_at,omitempty"` // 删除时间
	IsDel     bool       `json:"is_del" gorethink:"is_del,omitempty"`         // 是否删除
}

type Message struct {
	*Model
	Sender   string        `json:"sender" gorethink:"sender"`     // 发送者
	Receiver string        `json:"receiver" gorethink:"receiver"` // 接收者
	Content  string        `json:"content" gorethink:"content"`   // 消息内容
	State    MessageStatus `json:"state" gorethink:"state"`       // 消息状态
	Type     MessageType   `json:"type" gorethink:"type"`         // 消息类型
	// FileName string        `json:"file_name" gorethink:"file_name"` // 文件名称
	// Timestamp time.Time     `json:"timestamp" gorethink:"timestamp"` // 时间戳
}

func main() {
	// 连接WebSocket服务器
	// token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjU1NDdhYmU0LTM4NmUtNDVjMC1hZjBhLWFkOWNjOWI4MmU1OSJ9.01bbtcIqN5tSvnPCZ9Y_dPiuaq7qIWBILvRaZrCFWAY"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImU4MGMwNTUyLTMwMjctNDE0Yy1iMjA4LTQwMzFmZTczMGYzZCJ9.jl8pyHtchgp6kmaVhv4GP5V-PtIvhzKaBnRbGjIOXuY"
	header := ws.HandshakeHeaderHTTP{
		"Upgrade":               []string{"websocket"},
		"Connection":            []string{"Upgrade"},
		"Sec-WebSocket-Key":     []string{"dGhlIHNhbXBsZSBub25jZQ=="},
		"Sec-WebSocket-Version": []string{"13"},
		"Authorization":         []string{"Bearer " + token},
	}
	d := ws.Dialer{
		Header: header,
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	conn, _, _, err := d.Dial(ctx, "ws://127.0.0.1:8080/v1/ws")
	if err != nil {
		log.Println(err)
		return
	}
	// 接收消息
	go receiveMessages(conn)

	fmt.Println("连接成功！")

	// 从标准输入读入消息并发送到服务器
	scanner := bufio.NewScanner(os.Stdin)
	var msg Message
	for scanner.Scan() {
		msg = Message{
			Sender:   "e80c0552-3027-414c-b208-4031fe730f3d",
			Receiver: "5547abe4-386e-45c0-af0a-ad9cc9b82e59",
			Content:  "hello",
		}
		msg.Content = string(scanner.Bytes())
		data, err := json.Marshal(msg)
		if err != nil {
			log.Fatal("marshal message: %w", err)
		}
		err = wsutil.WriteClientMessage(conn, ws.OpText, data)
		if err != nil {
			log.Fatal(err)
			break
		}
	}
}

func receiveMessages(conn net.Conn) {
	for {
		data, op, err := wsutil.ReadServerData(conn)
		if err != nil {
			log.Println("read data: ", err)
			break
		}
		if op == ws.OpClose {
			conn.Close()
			log.Fatal("server close")
		}
		var msg Message
		json.Unmarshal(data, &msg)
		fmt.Printf("receive [%s] from [%s]\n", msg.Content, msg.Sender)
	}
}
