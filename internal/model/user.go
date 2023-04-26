// 说明：用户模型 定义用户结构体及其相关方法(验证、注册、登录)
package model

import (
	"time"

	"github.com/google/uuid"
)

// Define Gender and UserRole
type Gender int

const (
	Male    Gender = iota // 男性(默认)
	Female                // 女性
	Ladyboy               // 中性
	Other                 // 其他(未知或者隐藏)
)

type UserRole int

const (
	GeneralUser UserRole = iota // 普通人(默认)
	Admin                       // 管理员
	Bot                         // 机器人
)

type User struct {
	*Model
	Name           string            `json:"name,omitempty" gorethink:"name"`
	Email          string            `json:"email,omitempty" gorethink:"email,index"`
	Password       string            `json:"password,omitempty" gorethink:"password"`
	Gender         Gender            `json:"gender,omitempty" gorethink:"gender,omitempty"`
	Friends        map[string]string `json:"friends,omitempty" gorethink:"friends,omitempty"` // 好友列表 备注->UserID
	Groups         map[string]string `json:"groups,omitempty" gorethink:"groups,omitempty"`   // 所属群组ID 备注->GROUPID
	ProfileImage   string            `json:"profile_image,omitempty" gorethink:"profile_image,omitempty"`
	Role           UserRole          `json:"role,omitempty" gorethink:"role,omitempty"`
	Address        string            `json:"address,omitempty" gorethink:"address,omitempty"`
	Phone          string            `json:"phone,omitempty" gorethink:"phone,omitempty"`
	Birthdate      *time.Time        `json:"birthdate,omitempty" gorethink:"birthdate,omitempty"`
	SocialAccounts map[string]string `json:"social_accounts,omitempty" gorethink:"social_accounts,omitempty"`
	IsVerified     bool              `json:"is_verified,omitempty" gorethink:"is_verified,omitempty"` // 是否验证
	IsActive       bool              `json:"is_active,omitempty" gorethink:"is_active,omitempty"`     // 是否激活
	IsOnline       bool              `json:"is_online" gorethink:"is_online,omitempty"`               // 是否在线
	CreatedIP      string            `json:"created_ip,omitempty" gorethink:"created_ip,omitempty"`   // 创建 IP 地址
	UpdatedIP      string            `json:"updated_ip,omitempty" gorethink:"updated_ip,omitempty"`   // 更新 IP 地址
	LastLoginIP    string            `json:"last_login_ip,omitempty" gorethink:"last_login_ip,omitempty"`
	LastLoginAt    *time.Time        `json:"last_login_at,omitempty" gorethink:"last_login_at,omitempty"`
}

// NewUser 实例化User
func NewUser() *User {
	now := time.Now()
	return &User{
		Model: &Model{
			ID:        uuid.New().String(),
			CreatedAt: &now,
		},
	}
}
