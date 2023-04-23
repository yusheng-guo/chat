package dao

import (
	"fmt"

	"github.com/yushengguo557/chat/internal/model"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

// InsertUser 增加用户数据
func (d *Dao) InsertUser(user *model.User) error {
	_, err := r.Table("users").Insert(user).Run(d.Session)
	return fmt.Errorf("insert a user into the users table: %w", err)
}

// DeleteUserByID 通过 ID 删除用户
func (d *Dao) DeleteUserByID(uid string) error {
	_, err := r.Table("users").Filter(r.Row.Field("id").Eq(uid)).Delete().RunWrite(d.Session)
	return fmt.Errorf("delete user from the users table: %w", err)
}

// UpdateUserByID 通过 ID 更新用户
func (d *Dao) UpdateUserByID(uid string, data map[string]any) error {
	_, err := r.Table("users").Filter(r.Row.Field("id").Eq(uid).Update(data)).RunWrite(d.Session)
	return fmt.Errorf("update user information by user ID: %w", err)
}

// FindUserByID 通过 ID 查找用户
func (d *Dao) FindUserByID(uid string) (*model.User, error) {
	cursor, err := r.Table("users").Filter(r.Row.Field("id").Eq(uid)).Run(d.Session)
	if err != nil {
		return nil, fmt.Errorf("find user by user ID: %w", err)
	}
	var u *model.User
	err = cursor.One(u)
	return u, fmt.Errorf("get a user from the found by id: %w", err)
}

// FindUserByID 通过 name 查找用户
func (d *Dao) FindUserByName(name string) ([]*model.User, error) {
	cursor, err := r.Table("users").Filter(r.Row.Field("name").Eq(name)).Run(d.Session)
	if err != nil {
		return nil, fmt.Errorf("find user by user name: %w", err)
	}
	var users []*model.User
	err = cursor.All(users)
	return users, fmt.Errorf("get a user from the found by name: %w", err)
}
