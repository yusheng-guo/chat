package v1

import (
	"github.com/gin-gonic/gin"
)

// @Summary 添加好友
// @Description 当前用户添加好友
// @Tags friend
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param user_id path string true "目标用户 ID"
// @Success 200 {object} model.User
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /v1/friend/:id [post]
func AddFriend(c *gin.Context) {}

// @Summary 移除好友
// @Description 当前用户移除好友
// @Tags friend
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param user_id path string true "目标用户 ID"
// @Success 200 {object} model.User
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /v1/friend/:id [delete]
func DeleteFriend(c *gin.Context) {}

// @Summary 更新朋友信息
// @Description 当前用户更新好友信息
// @Tags friend
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param user_id path string true "目标用户 ID"
// @Success 200 {object} model.User
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /v1/friend/:id [put]
func UpdateFriendInfo(c *gin.Context) {}

// @Summary 获取朋友信息
// @Description 获取当前用户指定好友信息
// @Tags friend
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param user_id path string true "目标用户 ID"
// @Success 200 {object} model.User
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /v1/friend/:id [get]
func GetFriendInfo(c *gin.Context) {}

// @Summary 获取朋友列表
// @Description 获取当前用户所有好友
// @Tags friend
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {array} model.User
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /v1/friends/ [get]
func GetMyFriends(c *gin.Context) {}
