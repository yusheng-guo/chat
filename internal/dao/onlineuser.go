package dao

import (
	"context"
	"fmt"

	"github.com/yushengguo557/chat/internal/model"
)

var ctx = context.Background()

// ------------------------------------- 增加 -------------------------------------
// AddOnlineUser 添加在线用户
func (d *Dao) AddOnlineUser(user *model.OnlineUser) (err error) {
	// 1.序列化用户
	var serialized []byte
	if serialized, err = user.Marshal(); err != nil {
		return fmt.Errorf("serialize user struct when adding online user: %w", err)
	}

	// 2.添加user到redis中
	if err = d.Client.HSet(ctx, "online_users", user.ID, string(serialized)).Err(); err != nil {
		return fmt.Errorf("adding online user to redis: %w", err)
	}
	return nil
}

// ------------------------------------- 删除 -------------------------------------
// RemoveOnlineUser 移除在线用户
func (d *Dao) RemoveOnlineUser(id string) (err error) {
	// 1.找到用户 关闭 websocket 连接
	user, err := d.GetOnlineUser(id)
	if err != nil {
		return fmt.Errorf("get online user in `RemoveOnlineUser Function`: %w", err)
	}
	if err = user.Conn.Close(); err != nil {
		return fmt.Errorf("close WebSocket Connection: %w", err)
	}
	// 2.删除用户
	if err = d.Client.HDel(ctx, "online_users", id).Err(); err != nil {
		return fmt.Errorf("removing online user to redis: %w", err)
	}
	return nil
}

// ------------------------------------- 查询 -------------------------------------
func (d *Dao) GetOnlineUser(id string) (*model.OnlineUser, error) {
	// 1.从redis中获取
	serialized, err := d.Client.HGet(ctx, "online_users", id).Result()
	fmt.Println("serialized", serialized)
	if err != nil {
		return nil, fmt.Errorf("getting online user in redis database: %w", err)
	}
	// 2.反序列化
	var user model.OnlineUser
	user.Unmarshal([]byte(serialized))
	return &user, nil
}

// ------------------------------------- 其他 -------------------------------------
// IsOnline 用户是否在线
func (d *Dao) IsOnline(id string) bool {
	return d.Client.HExists(ctx, "online_users", id).Val()
}
