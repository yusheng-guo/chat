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
}

// New 实例化 Service
func New(ctx context.Context) Service {
	return Service{
		ctx: ctx,
		dao: dao.New(global.Session),
	}
}
