// 说明：用户模型 定义用户结构体及其相关方法(验证、注册、登录)
package model

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

// Define Gender and UserRole
type Gender int

const (
	Male    Gender = iota // 男性
	Female                // 女性
	Ladyboy               // 中性
	Other                 // 其他（未知）
)

type UserRole int

const (
	Admin       UserRole = iota // 管理员
	GeneralUser                 // 普通人
	Bot                         // 机器人
)

type User struct {
	*Model
	Name           string            `json:"name" gorethink:"name"`
	Email          string            `json:"email,omitempty" gorethink:"email,index"`
	Password       string            `json:"password,omitempty" gorethink:"password"`
	Gender         Gender            `json:"gender,omitempty" gorethink:"gender"`
	Friends        map[string]string `json:"friends,omitempty" gorethink:"friends"` // 好友列表 备注->UserID
	Groups         map[string]string `json:"groups,omitempty" gorethink:"groups"`   // 所属群组ID 备注->GROUPID
	ProfileImage   string            `json:"profile_image,omitempty" gorethink:"profile_image"`
	Role           UserRole          `json:"role,omitempty" gorethink:"role"`
	Address        string            `json:"address,omitempty" gorethink:"address"`
	Phone          string            `json:"phone,omitempty" gorethink:"phone"`
	Birthdate      *time.Time        `json:"birthdate,omitempty" gorethink:"birthdate"`
	SocialAccounts map[string]string `json:"social_accounts,omitempty" gorethink:"social_accounts"`
	IsVerified     bool              `json:"is_verified,omitempty" gorethink:"is_verified"` // 是否验证
	IsActive       bool              `json:"is_active,omitempty" gorethink:"is_active"`     // 是否激活
	IsOnline       bool              `json:"is_online" gorethink:"is_online"`               // 是否在线
	CreatedIP      string            `json:"created_ip,omitempty" gorethink:"created_ip"`   // 创建 IP 地址
	UpdatedIP      string            `json:"updated_ip,omitempty" gorethink:"updated_ip"`   // 更新 IP 地址
	LastLoginIP    string            `json:"last_login_ip,omitempty" gorethink:"last_login_ip"`
	LastLoginAt    *time.Time        `json:"last_login_at,omitempty" gorethink:"last_login_at"`
}

// 创建用户
// func NewUser() (u *User) {
// 	return
// }

// Register 用户注册
func RegisterUser(email, password string) error {
	// 如果数据库中有该对应关系 返回注册失败
	id := uuid.New().String()
	now := time.Now()
	u := &User{
		Model: &Model{
			ID:        id,
			CreatedAt: &now,
			UpdatedAt: nil,
			DeletedAt: nil,
			IsDel:     false,
		},
		Name:           email[:strings.IndexRune(email, '@')+1], // 默认用户名
		Email:          email,
		Password:       password,
		Friends:        nil,
		Groups:         nil,
		Address:        "",
		Phone:          "",
		Gender:         Other,       // 默认其他
		Birthdate:      nil,         // 未知
		ProfileImage:   "",          //未知
		SocialAccounts: nil,         // 暂时没有
		Role:           GeneralUser, // 默认普通用户
		IsVerified:     false,       // 没有验证
		IsActive:       true,        // 已激活
		IsOnline:       true,        // 在线 (与服务端建立连接)
		CreatedIP:      "",          // 当前连接的客户端地址
		UpdatedIP:      "",          //
		LastLoginIP:    "",          //
		LastLoginAt:    nil,         //
	}
	u.Password = password
	// 将用户信息保存到数据库

	return nil
}

// Login 用户登录
func LoginUser(email, password string) error {
	// TODO: 实现用户登录的逻辑
	// 查询数据库中 是否存在该对应关系
	// 如果没有该对应关系 登录失败
	return nil
}

// SendMessage 发送消息
func (u *User) SendMessage(content string, recipientID string) error {
	// TODO: 实现发送消息的逻辑
	return nil
}

// 用户信息管理

// Update 更新用户单个信息
func (u *User) Update(k, v string) error {

	return nil
}

// UpdateMulti 更新用户多个消息
func (u *User) UpdateMulti() error {

	return nil
}

// 朋友管理

// GetFriends 获取朋友列表
func (u *User) GetFriends() (*User, error) {
	// 查询数据库
	// 从 u.Friends 中获取用户ID
	// 根据用户ID 查询数据库 并返回
	return nil, nil
}

func (u *User) AddFriend(fname, fid string) error {
	// u.Friends = append(u.Friends, fid)
	u.Friends[fname] = fid
	// 保存到数据库
	return nil
}
