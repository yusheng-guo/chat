package dao

import (
	"github.com/redis/go-redis/v9"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type Dao struct {
	*r.Session    // rethink 数据库 会话
	*redis.Client // redis 数据库 客户端
}

// New 实例化 Dao 结构体
func NewDao(s *r.Session, c *redis.Client) *Dao {
	return &Dao{
		Session: s,
		Client:  c,
	}
}
