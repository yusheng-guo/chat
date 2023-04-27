package service

import (
	"context"

	"github.com/yushengguo557/chat/global"
	"github.com/yushengguo557/chat/internal/dao"
)

// Service
type Service struct {
	ctx context.Context
	dao *dao.Dao
	// conn *net.Conn
}

// New 实例化 Service
func NewService(ctx context.Context) Service {
	return Service{
		ctx: ctx,
		dao: dao.NewDao(global.Session, global.RedisClient),
	}
}
