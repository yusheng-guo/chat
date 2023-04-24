package dao

import (
	"fmt"

	"github.com/yushengguo557/chat/internal/model"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

// InsertUser 增加用户数据
func (d *Dao) InsertUser(user *model.User) error {
	_, err := r.DB("chat").Table("users").Insert(user).Run(d.Session)
	if err != nil {
		return fmt.Errorf("insert a user into the users table: %w", err)
	}
	return nil
}

// DeleteUserByID 通过 ID 删除用户
func (d *Dao) DeleteUserByID(uid string) error {
	_, err := r.DB("chat").Table("users").Filter(r.Row.Field("id").Eq(uid)).Delete().RunWrite(d.Session)
	if err != nil {
		return fmt.Errorf("delete user from the users table: %w", err)
	}
	return nil
}

// UpdateUserByID 通过 ID 更新用户
func (d *Dao) UpdateUserByID(uid string, data map[string]any) error {
	_, err := r.DB("chat").Table("users").Filter(r.Row.Field("id").Eq(uid).Update(data)).RunWrite(d.Session)
	if err != nil {
		return fmt.Errorf("update user information by user ID: %w", err)
	}
	return nil
}

// FindUserByID 通过 ID 查找用户
func (d *Dao) FindUserByID(uid string) (*model.User, error) {
	cursor, err := r.DB("chat").Table("users").Filter(r.Row.Field("id").Eq(uid)).Run(d.Session)
	if err != nil {
		return nil, fmt.Errorf("find user by user ID: %w", err)
	}
	u := new(model.User)
	err = cursor.One(u)
	if err != nil {
		return nil, fmt.Errorf("get a user from the found by id: %w", err)
	}
	return u, nil
}

// FindUsersByName 通过 name 查找用户
func (d *Dao) FindUsersByName(name string) ([]*model.User, error) {
	cursor, err := r.DB("chat").Table("users").Filter(r.Row.Field("name").Eq(name)).Run(d.Session)
	if err != nil {
		return nil, fmt.Errorf("find user by user name: %w", err)
	}
	users := []*model.User{}
	err = cursor.All(users)
	if err != nil {
		return nil, fmt.Errorf("get a user from the found by name: %w", err)
	}
	return users, nil
}

// FindUserByEmail 通过 name 查找用户
func (d *Dao) FindUserByEmail(email string) (*model.User, error) {
	cursor, err := r.DB("chat").Table("users").Filter(r.Row.Field("email").Eq(email)).Run(d.Session)
	if err != nil {
		return nil, fmt.Errorf("find user by user email: %w", err)
	}
	user := new(model.User)
	err = cursor.One(user)
	if err != nil {
		return nil, fmt.Errorf("get a user from the found by email: %w", err)
	}
	return user, nil
}
