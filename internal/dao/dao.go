package dao

import r "gopkg.in/rethinkdb/rethinkdb-go.v6"

type Dao struct {
	*r.Session
}

// New 实例化 Dao 结构体
func New(s *r.Session) *Dao {
	return &Dao{Session: s}
}
