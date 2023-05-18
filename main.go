package main

import (
	"fmt"
	"log"
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
		global.Logger.Warn(err)
	}
	global.Logger.Info("Database and tables created!")

	// 初始化配置
	err = global.InitConfig()
	if err != nil {
		// global.Logger.Warn(err)
		panic(err)
	}
	fmt.Println(global.Email.Host)
	fmt.Println(global.Email.Port)
	fmt.Println(global.Email.Username)
	fmt.Println(global.Email.Password)
	global.Logger.Info("Configuration file is loaded!")
}

// @title Chat System
// @version 1.0.0
// @description Chat system developed using Gin and github.com/gobwas/ws!
func main() {
	// 1.关闭数据库
	defer global.Session.Close()     // 关闭rethink数据库会话
	defer global.RedisClient.Close() // 关闭redis数据库客户端

	// 2.路由
	router := routers.NewRouter()
	// gin.DefaultWriter = colorable.NewColorableStdout() // windows 平台支持着色

	// 3.配置服务
	server := &http.Server{
		Addr:           ":" + global.ServerConfig.Port,
		Handler:        router,
		ReadTimeout:    global.ServerConfig.ReadTimeOut,
		WriteTimeout:   global.ServerConfig.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("listen at: %s\n", global.ServerConfig.Port)

	// 4.启动服务 (监听)
	if err := server.ListenAndServe(); err != nil {
		global.Logger.Warn(err)
	}
}
