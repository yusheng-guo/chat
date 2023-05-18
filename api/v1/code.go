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
	fmt.Println("creat new email service")
	e, err := utils.NewEmailService()
	if err != nil {
		panic(err)
	}

	// 关闭 客户端
	// defer func() {
	// 	_ = e.Client.Close()
	// }()

	// 发送 验证码
	fmt.Println("send email captcha")
	err = e.SendCaptcha("declinedyew@outlook.com")
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"meg": "OK"})
}
