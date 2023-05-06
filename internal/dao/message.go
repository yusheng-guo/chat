package dao

import (
	"fmt"
	"time"

	"github.com/yushengguo557/chat/internal/model"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

// ------------------------------------- 增加 -------------------------------------
// Insertmessage 增加数据
func (d *Dao) InsertMessage(msg *model.Message) error {
	_, err := r.DB("chat").Table("messages").
		Insert(msg).
		RunWrite(d.Session)
	if err != nil {
		return fmt.Errorf("insert a message into the messages table: %w", err)
	}
	return nil
}

// ------------------------------------- 删除 -------------------------------------
// DeleteMessageByID 通过消息ID	删除消息 (底层：将消息 is_del 字段标记为 true)
func (d *Dao) DeleteMessageByID(mid string) error {
	return d.UpdatemessageByID(mid,
		map[string]any{"is_del": true})
}

// DeleteMessageByID 通过 ID 删除消息 (底层：从rethinkdb中删除消息)
func (d *Dao) RealDeleteMessageByID(mid string) error {
	_, err := r.DB("chat").Table("messages").
		Filter(r.Row.Field("id").Eq(mid)).
		Delete().
		RunWrite(d.Session)
	if err != nil {
		return fmt.Errorf("delete message from the messages table: %w", err)
	}
	return nil
}

// ------------------------------------- 更新 -------------------------------------
// UpdatemessageByID 通过 ID 更新消息信息
func (d *Dao) UpdatemessageByID(mid string, data map[string]any) error {
	data["updated_at"] = time.Now() // 更新时间
	_, err := r.DB("chat").Table("messages").
		Filter(r.Row.Field("id").Eq(mid)).
		Update(data).
		RunWrite(d.Session)
	// _, err := r.DB("chat").Table("messages").
	// 	Get(mid).
	// 	Update(data).
	// 	RunWrite(d.Session)
	if err != nil {
		return fmt.Errorf("update message information by message ID: %w", err)
	}
	return nil
}

// UpdatemessageContentByID 通过消息ID更新消息内容
func (d *Dao) UpdatemessageContentByID(mid, content string) error {
	return d.UpdatemessageByID(mid,
		map[string]any{"content": content})
}

// UpdatemessageStatueByID 通过消息ID更新消息状态
func (d *Dao) UpdatemessageStatueByID(mid string, state model.MessageStatus) error {
	return d.UpdatemessageByID(mid,
		map[string]any{"state": state})
}

// ------------------------------------- 查找 -------------------------------------
// FindmessageByID 通过 ID 查找消息
func (d *Dao) FindmessageByID(mid string) (*model.Message, error) {
	cursor, err := r.DB("chat").Table("messages").
		Filter(r.Row.Field("id").Eq(mid)).
		Run(d.Session)
	if err != nil {
		return nil, fmt.Errorf("find messages by message ID: %w", err)
	}
	var m *model.Message
	err = cursor.One(m)
	if err != nil {
		return nil, fmt.Errorf("get a message from the found: %w", err)
	}
	return m, nil
}
