// 说明：全局变量
package global

import (
	"github.com/yushengguo557/chat/config"
)

var (
	DBConfig     *config.DBConfigS     // 数据库配置
	ServerConfig *config.ServerConfigS // 服务器配置
)

func InitConfig() (err error) {
	// TODO: 建立配置 对全局变量进行赋值
	var cfg *config.Config
	cfg, err = config.NewConfig()
	if err != nil {
		return err
	}
	err = cfg.ReadSection("Server", &ServerConfig)
	if err != nil {
		return err
	}
	err = cfg.ReadSection("Database", &DBConfig)
	if err != nil {
		return err
	}
	return
}
