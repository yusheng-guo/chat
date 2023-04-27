package v1

import (
	"github.com/gin-gonic/gin"
)

// @Summary 发送消息
// @Description 发送我要发送的消息
// @Tags message
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /v1/message [post]
func SendMessage(c *gin.Context) {
}

// @Summary 删除消息
// @Description 删除我发送的消息
// @Tags message
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param user_id path string true "目标消息 ID"
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /v1/message/:id [delete]
func DeleteMessage(c *gin.Context) {
}

// @Summary 更新消息
// @Description 更新我发送的消息
// @Tags message
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param user_id path string true "目标消息 ID"
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /v1/message/:id [put]
func UpdateMessage(c *gin.Context) {
}

// @Summary 接收消息
// @Description 接收发送给我的消息
// @Tags message
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /v1/message [get]
func ReceiveMessage(c *gin.Context) {
	// svc := service.NewService(c.Request.Context())
}
