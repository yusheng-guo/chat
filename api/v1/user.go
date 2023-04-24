package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yushengguo557/chat/internal/service"
)

// @Summary 用户注册
// @Description 使用电子邮件进行注册
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
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
	err := svc.Register(param) // 用户注册
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User register successfully"})
}

// @Summary 用户登录
// @Description 使用电子邮件和明码进行登录
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /v1/login [post]
func Login(c *gin.Context) {}

// @Summary 用户登出
// @Description Log out
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
// @Description 管理员登录
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
// @Description 获取我的个人信息
// @Tags me
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /v1/me/info [get]
func GetMyInfo(c *gin.Context) {}

// @Summary 更新我的个人信息
// @Description 更新我的个人信息
// @Tags me
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /v1/me/info [put]
func UpdateMyInfo(c *gin.Context) {}
