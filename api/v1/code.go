package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yushengguo557/chat/utils"
)

// 获取邮箱验证码 测试
func GetCode(c *gin.Context) {
	// 新建 Email服务
	e, err := utils.NewEmailService()
	if err != nil {
		fmt.Println("new... ", err)
		panic(fmt.Errorf("new...%w", err))
	}

	// 关闭 客户端
	defer func() {
		fmt.Println("close...")
		_ = e.Client.Close()
	}()

	// err := e.Send("declinedyew@outlook.com", "验证码", GenRandFourDigits()())
	// 发送 验证码
	err = e.SendCaptcha("declinedyew@outlook.com")
	if err != nil {
		fmt.Println("send... ", err)
		panic(fmt.Errorf("send... %w", err))
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
