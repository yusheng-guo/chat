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
	"github.com/yushengguo557/chat/client/message"
)

func main() {
	// 连接WebSocket服务器
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
	var msg message.Message
	for scanner.Scan() {
		msg = message.Message{
			Sender:    "e80c0552-3027-414c-b208-4031fe730f3d",
			Receiver:  "5547abe4-386e-45c0-af0a-ad9cc9b82e59",
			Content:   "hello",
			State:     message.MessageStatusSent,
			Type:      message.MessageTypeText,
			CreatedAt: int(time.Now().UnixMilli()),
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
		var msg message.Message
		json.Unmarshal(data, &msg)
		fmt.Printf("receive [%s] from [%s]\n", msg.Content, msg.Sender)
	}
}
