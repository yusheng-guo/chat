package main

import (
	"net/http"

	"github.com/yushengguo557/chat/global"
	"github.com/yushengguo557/chat/internal/routers"
)

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
}
