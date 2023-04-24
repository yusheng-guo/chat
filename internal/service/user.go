// 说明：user相关业务逻辑
package service

import (
	"fmt"
	"strings"

	"github.com/yushengguo557/chat/internal/model"
)

// RegisterRequest 注册请求
type RegisterRequest struct {
	Email    string `form:"email"`    // 邮箱
	Password string `form:"password"` // 密码
}

// Register 注册用户
func (s *Service) Register(r *RegisterRequest) (err error) {
	// 1.验证邮箱\密码合法性
	// 2.查询数据库中是否已经有该邮箱
	var user *model.User
	user, _ = s.dao.FindUserByEmail(r.Email)
	if user != nil {
		return fmt.Errorf("users are already registered")
	}
	// 3.创建用户
	user = model.NewUser()
	user.Email = r.Email
	user.Password = r.Password
	user.Name = r.Email[:strings.IndexRune(r.Email, '@')] // 默认用户名
	// 4.将用户信息保存到数据库
	err = s.dao.InsertUser(user)
	if err != nil {
		return fmt.Errorf("insert user when registering a user: %w", err)
	}
	return nil
}
