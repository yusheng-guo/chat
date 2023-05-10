package v1

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yushengguo557/chat/api/common"
	"github.com/yushengguo557/chat/internal/service"
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
func AddFriend(c *gin.Context) {
	myid, exists := c.Get("id")
	if exists {
		log.Panic("id not exists")
	}
	friendid := c.PostForm("id")
	svc := service.NewService(c)
	err := svc.AddFriendByID(myid.(string), friendid)
	if err != nil {
		rsp := common.NewResponse(common.InternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, rsp)
		return
	}
	rsp := common.NewResponse(common.OK, "success")
	c.JSON(http.StatusOK, rsp)
}

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
func DeleteFriend(c *gin.Context) {
	myid, exists := c.Get("id")
	if exists {
		log.Panic("id not exists")
	}
	friendid := c.PostForm("id")
	svc := service.NewService(c)
	err := svc.DeleteFriendByID(myid.(string), friendid)
	if err != nil {
		rsp := common.NewResponse(common.InternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, rsp)
		return
	}
	rsp := common.NewResponse(common.OK, "success")
	c.JSON(http.StatusOK, rsp)
}

// @Summary 更新朋友备注
// @Description 当前用户更新好友备注
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
func UpdateFriendNote(c *gin.Context) {
	// 从 context 中获取自身 id
	myid, exists := c.Get("id")
	if exists {
		log.Panic("id not exists")
	}

	// 从put请求中获取备注
	var mfnr common.ModifyFriendNoteRequest
	err := c.BindJSON(&mfnr)
	if err != nil {
		rsp := common.NewResponse(common.InternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, rsp)
		return
	}
	note := mfnr.Note

	// 获取 朋友id
	friendid := c.Param("id")

	// 修改备注
	svc := service.NewService(c)
	err = svc.ModifyFriendNoteByID(myid.(string), friendid, note)
	if err != nil {
		rsp := common.NewResponse(common.InternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, rsp)
		return
	}

	// 响应 数据
	rsp := common.NewResponse(common.OK, "success")
	c.JSON(http.StatusOK, rsp)
}

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
func GetFriendInfo(c *gin.Context) {
	// 从 context 中获取自身 id
	myid, exists := c.Get("id")
	if exists {
		log.Panic("id not exists")
	}

	// 获取我的所有好友
	friendid := c.Param("id")
	svc := service.NewService(c)
	friend, err := svc.GetFriendInfoByID(myid.(string), friendid)
	if err != nil {
		rsp := common.NewResponse(common.InternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, rsp)
		return
	}

	// 响应 数据
	rsp := common.NewResponse(common.OK, "success")
	rsp.Data = friend
	c.JSON(http.StatusOK, rsp)
}

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
func GetMyFriends(c *gin.Context) {
	// 从 context 中获取自身 id
	myid, exists := c.Get("id")
	if exists {
		log.Panic("id not exists")
	}

	// 获取我的所有好友
	svc := service.NewService(c)
	friends, err := svc.GetMyFriends(myid.(string))
	if err != nil {
		rsp := common.NewResponse(common.InternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, rsp)
		return
	}

	// 响应 数据
	rsp := common.NewResponse(common.OK, "success")
	rsp.Data = friends
	c.JSON(http.StatusOK, rsp)
}
