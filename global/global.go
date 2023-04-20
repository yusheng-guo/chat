// 说明：全局变量
package global

import (
	"log"

	"github.com/yushengguo557/chat/config"
)

var (
	DBConfig     *config.DBConfigS     // 数据库配置
	ServerConfig *config.ServerConfigS // 服务器配置
)

func init() {
	// TODO: 建立配置 对全局变量进行赋值
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = cfg.ReadSection("Server", &ServerConfig)
	if err != nil {
		log.Fatal(err)
	}
	err = cfg.ReadSection("Database", &DBConfig)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Configuration file is loaded!")
}
