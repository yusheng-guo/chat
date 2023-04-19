package main

import (
	"net/http"
	"time"

	"github.com/yushengguo557/chat/internal/routers"
)

func main() {
	r := routers.NewRouter()
	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
