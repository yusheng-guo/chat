package global

import (
	"context"
	"fmt"

	redis "github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

// InitRedisDB 初始化 redis 数据库
func InitRedisDB() (err error) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         "119.91.204.226:6379",
		Password:     "",     // no password set
		DB:           0,      // use default DB
		PoolSize:     50,     // 最大连接数
		MinIdleConns: 10,     // 闲置连接的最小数量
		MaxIdleConns: 20,     // 闲置连接的最大数量
		PoolTimeout:  60 * 3, // 连接池可用连接等待超时时间(秒)
	})
	// ping 测试
	_, err = RedisClient.Ping(context.Background()).Result()
	if err != nil {
		return fmt.Errorf("ping redis database %w", err)
	}
	return
}
