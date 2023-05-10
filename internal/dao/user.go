package dao

import (
	"fmt"
	"log"
	"time"

	"github.com/yushengguo557/chat/internal/model"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

// ------------------------------------- 增加 -------------------------------------
// InsertUser 增加用户数据
func (d *Dao) InsertUser(user *model.User) error {
	_, err := r.DB("chat").Table("users").
		Insert(user).
		RunWrite(d.Session)
	if err != nil {
		return fmt.Errorf("insert a user into the users table: %w", err)
	}
	return nil
}

// ------------------------------------- 删除 -------------------------------------
// DeleteUserByID 通过 ID 删除用户 (底层：将用户 字段 is_del 标记为 true)
func (d *Dao) DeleteUserByID(uid string) error {
	return d.UpdateUserByID(uid,
		&map[string]any{"id_del": true})
}

// RealDeleteUserByID 通过 ID 删除用户 (底层：将用户从rethinkdb中永久删除)
func (d *Dao) RealDeleteUserByID(uid string) error {
	_, err := r.DB("chat").Table("users").
		Filter(r.Row.Field("id").Eq(uid)).
		Delete().
		RunWrite(d.Session)
	if err != nil {
		return fmt.Errorf("delete user from the users table: %w", err)
	}
	return nil
}

// ------------------------------------- 更新 -------------------------------------
// UpdateUserByID 通过 ID 更新用户
func (d *Dao) UpdateUserByID(uid string, dataptr *map[string]any) error {
	data := *dataptr                // 解引用
	data["updated_at"] = time.Now() // 更新时间
	_, err := r.DB("chat").Table("users").
		Filter(r.Row.Field("id").Eq(uid)).
		Update(data).
		RunWrite(d.Session)
	if err != nil {
		return fmt.Errorf("update user information by user ID: %w", err)
	}
	return nil
}

// ------------------------------------- 查找 -------------------------------------
// FindUserByID 通过 ID 查找用户
func (d *Dao) FindUserByID(uid string) (*model.User, error) {
	cursor, err := r.DB("chat").Table("users").
		Filter(r.Row.Field("id").Eq(uid)).
		Run(d.Session)
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
	cursor, err := r.DB("chat").Table("users").
		Filter(r.Row.Field("name").Eq(name)).
		Run(d.Session)
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
	cursor, err := r.DB("chat").Table("users").
		Filter(r.Row.Field("email").Eq(email)).
		Run(d.Session)
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

// IsExist 判断用户是否存在
func (d *Dao) IstUserExis(uid string) bool {
	res, err := r.DB("chat").Table("users").
		Get(uid).
		Run(d.Session)
	if err != nil {
		log.Panic(fmt.Errorf("find user by user ID: %w", err))
		return false
	}
	if res.IsNil() {
		return false
	}
	return true
}
