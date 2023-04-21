package dao

import r "gopkg.in/rethinkdb/rethinkdb-go.v6"

type Dao struct {
	*r.Session
}

// NewDao 实例化 Dao 结构体
func NewDao(s *r.Session) *Dao {
	return &Dao{Session: s}
}
