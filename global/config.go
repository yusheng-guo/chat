// 说明：全局变量
package global

import (
	"fmt"

	"github.com/yushengguo557/chat/config"
)

var (
	DBConfig     *config.DBConfigS     // 数据库配置
	ServerConfig *config.ServerConfigS // 服务器配置
	Storage      *config.StorageS      // 存储
	Email        *config.EmailS        // 邮件
)

// InitConfig 建立配置 对全局变量进行赋值
func InitConfig() (err error) {
	var cfg *config.Config
	cfg, err = config.NewConfig()
	if err != nil {
		return fmt.Errorf("creating a new config: %w", err)
	}
	err = cfg.ReadSection("Server", &ServerConfig)
	if err != nil {
		return fmt.Errorf("reading section to global var ServerConfig: %w", err)
	}
	err = cfg.ReadSection("Database", &DBConfig)
	if err != nil {
		return fmt.Errorf("reading section to global var DBConfig: %w", err)
	}
	err = cfg.ReadSection("Storage", &Storage)
	if err != nil {
		return fmt.Errorf("reading section to global var Storage: %w", err)
	}
	err = cfg.ReadSection("Email", &Email)
	if err != nil {
		return fmt.Errorf("reading section to global var Email: %w", err)
	}
	return
}
