package dao

import (
	"fmt"

	"github.com/yushengguo557/chat/internal/model"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

// Insertmessage 增加数据
func (d *Dao) InsertMessage(msg *model.Message) error {
	_, err := r.DB("chat").Table("messages").Insert(msg).Run(d.Session)
	if err != nil {
		return fmt.Errorf("insert a message into the messages table: %w", err)
	}
	return nil
}

// DeleteMessageByID 通过 ID 删除消息
func (d *Dao) DeleteMessageByID(mid string) error {
	_, err := r.DB("chat").Table("messages").Filter(r.Row.Field("id").Eq(mid)).Delete().RunWrite(d.Session)
	if err != nil {
		return fmt.Errorf("delete message from the messages table: %w", err)
	}
	return nil
}

// UpdatemessageByID 通过 ID 更新消息
func (d *Dao) UpdatemessageByID(mid string, data map[string]any) error {
	_, err := r.DB("chat").Table("messages").Filter(r.Row.Field("id").Eq(mid).Update(data)).RunWrite(d.Session)
	if err != nil {
		return fmt.Errorf("update message information by message ID: %w", err)
	}
	return nil
}

// FindmessageByID 通过 ID 查找消息
func (d *Dao) FindmessageByID(mid string) (*model.Message, error) {
	cursor, err := r.DB("chat").Table("messages").Filter(r.Row.Field("id").Eq(mid)).Run(d.Session)
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
