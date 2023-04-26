package global

import "fmt"

// InitDB 初始化数据库
func InitDB() (err error) {
	err = InitRedisDB() // 初始化redisdb
	if err != nil {
		return fmt.Errorf("init redis database %w", err)
	}
	err = InitRethinkDB() // 初始化rethinkdb
	if err != nil {
		return fmt.Errorf("init rethinkdb %w", err)
	}
	return nil
}
