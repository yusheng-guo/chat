package v1

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yushengguo557/chat/api/common"
	"github.com/yushengguo557/chat/internal/model"
	"github.com/yushengguo557/chat/internal/service"
	"github.com/yushengguo557/chat/utils"
)

// @Summary 用户注册
// @Description User Registration
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /v1/register [post]
func Register(c *gin.Context) {
	param := &service.RegisterRequest{}
	svc := service.NewService(c.Request.Context())
	if err := c.Bind(param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ret := svc.Register(param) // 用户注册
	c.JSON(ret.Code, gin.H{"message": ret.Message})
}

// @Summary 用户登录
// @Description User Login
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /v1/login [post]
func Login(c *gin.Context) {
	// 1.登录参数
	param := &service.LoginRequest{}
	svc := service.NewService(c.Request.Context())
	if err := c.Bind(param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 2.携带参数进行登录
	u, err := svc.Login(param)
	if err != nil { // 登录失败
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 3.生成Token
	var token string
	token, err = utils.GenerateToken(u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// @Summary 用户登出
// @Description Log Out
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /v1/logout [post]
func Logout(c *gin.Context) {}

// @Summary 管理员登录
// @Description Administrator Login
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /v1/admin [post]
func Admin(c *gin.Context) {}

// @Summary 获取我的个人信息
// @Description Get my personal information.
// @Tags me
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /v1/me/info [get]
func GetMyInfo(c *gin.Context) {
	// 从 context 中获取自身 id
	myid, exists := c.Get("id")
	if exists {
		log.Panic("id not exists")
	}

	// 获取我的信息
	svc := service.NewService(c)
	me, err := svc.GetMyInfoByID(myid.(string))
	if err != nil {
		rsp := common.NewResponse(common.InternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, rsp)
		return
	}

	// 响应 数据
	rsp := common.NewResponse(common.OK, "success")
	rsp.Data = me
	c.JSON(http.StatusOK, rsp)
}

// @Summary 更新我的个人信息
// @Description Update my personal information.
// @Tags me
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /v1/me/info [put]
func UpdateMyInfo(c *gin.Context) {
	// 从 context 中获取自身 id
	myid, exists := c.Get("id")
	if exists {
		log.Panic("id not exists")
	}

	// 从put请求中获取备注
	var u model.User
	err := c.BindJSON(&u)
	if err != nil {
		rsp := common.NewResponse(common.InternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, rsp)
		return
	}

	// 修改信息
	svc := service.NewService(c)
	err = svc.ModifyMyInfoByID(myid.(string), &u)
	if err != nil {
		rsp := common.NewResponse(common.InternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, rsp)
		return
	}

	// 响应 数据
	rsp := common.NewResponse(common.OK, "success")
	c.JSON(http.StatusOK, rsp)

}
