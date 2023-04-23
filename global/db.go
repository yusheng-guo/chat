package global

import (
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

var Session *r.Session

func InitDB() (err error) {
	// TODO: 初始化数据库
	// Connection Pool
	// web url: http://119.91.204.226:32770/
	Session, err = r.Connect(r.ConnectOpts{
		Address:    "119.91.204.226:32769",
		InitialCap: 10,
		MaxOpen:    10,
	})
	if err != nil {
		return
	}
	// 创建 chat 数据库
	_, err = r.DBCreate("chat").RunWrite(Session)
	if err != nil {
		return
	}
	// 创建 users 表
	_, err = r.DB("chat").TableCreate("users").RunWrite(Session)
	if err != nil {
		return
	}
	// 创建 messages 表
	_, err = r.DB("chat").TableCreate("messages").RunWrite(Session)
	if err != nil {
		return
	}
	return
}
