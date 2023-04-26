package global

import (
	"fmt"

	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

var Session *r.Session

var ErrEmptyResult = r.ErrEmptyResult

// InitRethinkDB 初始化数据库
func InitRethinkDB() (err error) {
	// Connection Pool
	// web url: http://119.91.204.226:32770/
	// r.error
	Session, err = r.Connect(r.ConnectOpts{
		Address:    "119.91.204.226:32769",
		InitialCap: 10,
		MaxOpen:    10,
	})
	if err != nil {
		return fmt.Errorf("connecting to the rethinkdb: %w", err)
	}
	// 创建 chat 数据库
	_, err = r.DBCreate("chat").RunWrite(Session)
	if err != nil {
		return fmt.Errorf("creating database chat with rethinkdb: %w", err)
	}
	// 创建 users 表
	_, err = r.DB("chat").TableCreate("users").RunWrite(Session)
	if err != nil {
		return fmt.Errorf("creating table users in chat: %w", err)
	}
	// 创建 messages 表
	_, err = r.DB("chat").TableCreate("messages").RunWrite(Session)
	if err != nil {
		return fmt.Errorf("creating table messages in chat: %w", err)
	}
	return
}
