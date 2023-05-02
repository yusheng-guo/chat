package v1

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"github.com/yushengguo557/chat/global"
	"github.com/yushengguo557/chat/internal/service"
)

// @Summary 用户通信
// @Description 通过建立websocket协议进行用户间通信 收发消息
// @Tags ws
// @Accept json
// @Produce json
// @Param user_id path string true "ID"
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /v1/ws [get]
func HandleWebSocket(c *gin.Context) {
	// 1.从 Context 中获取 用户id
	id := c.GetString("id")
	fmt.Println("---id:", id, "---")
	// 2..升级为 websocket 连接
	conn, _, _, err := ws.UpgradeHTTP(c.Request, c.Writer)
	if err != nil {
		global.Logger.Warn("upgrading http to websocket: %w", err)
		c.JSON(http.StatusUpgradeRequired, gin.H{"message": "conn't upgrade http to websocket."})
	}
	fmt.Println("连接成功")
	// 3..关闭连接
	// defer conn.Close()

	// 4..将连接加入到全局变量OnlineUsers中
	global.OnlineUsers[id] = conn
	fmt.Println("添加在线用户")
	fmt.Println(conn.RemoteAddr().String())

	// 5.创建服务 进行通信
	svc := service.NewService(c.Request.Context())
	go func() {
		err := svc.Communicate(conn)
		log.Panicln(err)
	}()
}