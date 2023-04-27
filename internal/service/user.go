// 说明：user相关业务逻辑
package service

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/yushengguo557/chat/global"
	"github.com/yushengguo557/chat/internal/model"
)

// RegisterRequest 注册请求
type RegisterRequest struct {
	Email    string `form:"email"`    // 邮箱
	Password string `form:"password"` // 密码
}

// LoginRequest 登录请求
type LoginRequest struct {
	Email    string `form:"email"`    // 邮箱
	Password string `form:"password"` // 密码
}

// LoginResponse 登录响应
type RegisterResponse struct {
	Code    int    `form:"code"`    // 状态码
	Message string `form:"message"` // 响应消息
}

// LoginResponse 登录响应
type LoginResponse struct {
	Code    int    `form:"code"`    // 状态码
	Message string `form:"message"` // 响应消息
}

// Register 注册用户
func (s *Service) Register(r *RegisterRequest) *RegisterResponse {
	// 1.验证邮箱\密码合法性
	if !strings.Contains(r.Email, "@") {
		return &RegisterResponse{
			Code:    http.StatusBadRequest,
			Message: "Mailbox format error.",
		}
	}

	// 2.查询数据库中是否已经有该邮箱
	var user *model.User
	user, _ = s.dao.FindUserByEmail(r.Email)
	if user != nil {
		return &RegisterResponse{
			Code:    http.StatusConflict,
			Message: "User are already registered.",
		}
	}

	// 3.创建用户
	user = model.NewUser()
	user.Email = r.Email
	user.Password = r.Password
	user.Name = r.Email[:strings.IndexRune(r.Email, '@')] // 默认用户名

	// 4.将用户信息保存到数据库
	err := s.dao.InsertUser(user)
	if err != nil {
		global.Logger.Warn(fmt.Errorf("insert user when registering a user: %w", err))
		return &RegisterResponse{
			Code:    http.StatusInternalServerError,
			Message: "Try again later.",
		}
	}
	return &RegisterResponse{
		Code:    http.StatusOK,
		Message: "Register successfully",
	}
}

// Login 用户登录
func (s Service) Login(re *LoginRequest) *LoginResponse {
	// 1.该用户是否存在
	u, _ := s.dao.FindUserByEmail(re.Email)
	if u == nil {
		return &LoginResponse{
			Code:    http.StatusUnauthorized,
			Message: "User does not exist.",
		}
	}

	// 2.密码是否正确
	if u.Password != re.Password {
		return &LoginResponse{
			Code:    http.StatusUnauthorized,
			Message: "Password error.",
		}
	}
	// 3.新建在线用户
	onlineuser := model.NewOnlineUser(u.ID)

	// 4.加入到在线用户redis数据库中
	if err := s.dao.AddOnlineUser(onlineuser); err != nil {
		fmt.Errorf("adding onlineuser to redis in `Login Function`: %w", err)
		return &LoginResponse{
			Code:    http.StatusInternalServerError,
			Message: "Try again later.",
		}
	}

	// 5.响应
	return &LoginResponse{
		Code:    http.StatusOK,
		Message: "Login successful!",
	}
}
