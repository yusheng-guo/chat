package main

import (
	"net/http"

	"github.com/yushengguo557/chat/global"
	"github.com/yushengguo557/chat/internal/routers"
)

func init() {
	// 实例化 Logger
	global.InitLog()
	var err error

	// 初始化数据库 (忽略错误)
	_ = global.InitDB()
	if err != nil {
		global.Logger.Info(err)
	}
	global.Logger.Println("Database and tables created!")

	// 初始化配置
	err = global.InitConfig()
	if err != nil {
		global.Logger.Warn(err)
	}
	global.Logger.Info("Configuration file is loaded!")
}

func main() {
	router := routers.NewRouter()
	server := &http.Server{
		Addr:           ":" + global.ServerConfig.Port,
		Handler:        router,
		ReadTimeout:    global.ServerConfig.ReadTimeOut,
		WriteTimeout:   global.ServerConfig.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
	defer global.Session.Close() // 关闭数据库会话
}
