package main

import (
	"log"
	"net/http"
	"os"

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
		global.Logger.Warn(err)
	}
	global.Logger.Info("Database and tables created!")

	// 初始化配置
	err = global.InitConfig()
	if err != nil {
		global.Logger.Warn(err)
	}
	global.Logger.Info("Configuration file is loaded!")
}

// @title Chat System
// @version 1.0.0
// @description Chat system developed using Gin and github.com/gobwas/ws!
func main() {
	var err error
	router := routers.NewRouter()
	// gin.DefaultWriter = colorable.NewColorableStdout() // windows 平台支持着色
	server := &http.Server{
		Addr:           ":" + global.ServerConfig.Port,
		Handler:        router,
		ReadTimeout:    global.ServerConfig.ReadTimeOut,
		WriteTimeout:   global.ServerConfig.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("listen at: %s\n", global.ServerConfig.Port)
	if err = server.ListenAndServe(); err != nil {
		global.Session.Close() // 关闭数据库会话
		global.Logger.Warn(err)
		os.Exit(1)
	}
}
