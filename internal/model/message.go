// 说明：消息模型 创建消息 保存消息
package model

import (
	"time"

	"github.com/yushengguo557/chat/global"
	"github.com/yushengguo557/chat/utils"
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

// 消息定义
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

// 创建消息
func NewMessage(sender string, receiver string, content string) *Message {
	// 生成消息 ID
	id := utils.GenerateUuid()
	now := time.Now()
	// 创建消息并返回
	return &Message{
		Model: &Model{
			ID:        id,
			CreatedAt: &now,
			UpdatedAt: nil,
			DeletedAt: nil,
			IsDel:     false,
		},
		Sender:   sender,
		Receiver: receiver,
		Content:  content,
		State:    MessageStatusSent, // 默认值
		Type:     MessageTypeText,   // 默认值
	}
}

// 序列化消息
func (message *Message) Marshal() ([]byte, error) {
	return global.Json.Marshal(message)
}

// 反序列化
func (message *Message) Unmarshal(data []byte) error {
	return global.Json.Unmarshal(data, &message)
}

// -------------------------------------------------------------------------------
// TransferredMessage 用于 传输 的 消息
type TransferredMessage struct {
	ID        string         `json:"id"`         // 消息id
	Sender    string         `json:"sender"`     // 发送者
	Receiver  string         `json:"receiver"`   // 接收者
	Content   string         `json:"content"`    // 消息内容
	State     MessageStatus  `json:"state"`      // 消息状态
	Type      MessageType    `json:"type"`       // 消息类型
	CreatedAt *time.Duration `json:"created_at"` // 创建时间

}

// 创建消息
func NewTransferredMessage(sender string, receiver string, content string) *TransferredMessage {
	id := utils.GenerateUuid()
	return &TransferredMessage{
		ID:       id,
		Sender:   sender,
		Receiver: receiver,
		Content:  content,
		State:    MessageStatusSent, // 默认值
		Type:     MessageTypeText,   // 默认值
	}
}

// Message({
//     required this.id,
//     required this.sender,
//     required this.receiver,
//     required this.content,
//     required this.createdAt,
//     this.state = MessageStatus.messageStatusSent,
//     this.type = MessageType.messageTypeText,
//   });
