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
	// 验证邮箱合法性
	// 验证密码合法性
	// 查询数据库中是否已经有该邮箱
	var user *model.User
	user, err = s.dao.FindUserByEmail(r.Email)
	if err != nil {
		return err
	}
	if user != nil {
		return fmt.Errorf("users are already registered")
	}
	user = model.NewUser()
	user.Email = r.Email
	user.Password = r.Password
	user.Name = r.Email[:strings.IndexRune(r.Email, '@')+1] // 默认用户名

	// 将用户信息保存到数据库
	return s.dao.InsertUser(user)
}
