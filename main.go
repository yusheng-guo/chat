package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	var err error
	router := routers.NewRouter()
	// gin.DefaultWriter = colorable.NewColorableStdout() // windows 平台支持着色
	gin.DefaultWriter = global.Logger.Writer()
	server := &http.Server{
		Addr:           ":" + global.ServerConfig.Port,
		Handler:        router,
		ReadTimeout:    global.ServerConfig.ReadTimeOut,
		WriteTimeout:   global.ServerConfig.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}
	err = server.ListenAndServe()
	if err != nil {
		global.Logger.Warn(err)
	}
	defer global.Session.Close() // 关闭数据库会话
}
