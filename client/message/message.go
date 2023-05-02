package message

import "time"

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
