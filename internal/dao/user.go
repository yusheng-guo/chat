package dao

import (
	"github.com/yushengguo557/chat/internal/model"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

// InsertUser 增加用户数据
func (d *Dao) InsertUser(user *model.User) error {
	_, err := r.Table("users").Insert(user).Run(d.Session)
	return err
}

// DeleteUserByID 通过 ID 删除用户
func (d *Dao) DeleteUserByID(uid string) error {
	_, err := r.Table("users").Filter(r.Row.Field("id").Eq(uid)).Delete().RunWrite(d.Session)
	return err
}

// UpdateUserByID 通过 ID 更新用户
func (d *Dao) UpdateUserByID(uid string, data map[string]any) error {
	_, err := r.Table("users").Filter(r.Row.Field("id").Eq(uid).Update(data)).RunWrite(d.Session)
	return err
}

// FindUserByID 通过 ID 查找用户
func (d *Dao) FindUserByID(uid string) (u *model.User, err error) {
	cursor, err := r.Table("users").Filter(r.Row.Field("id").Eq(uid)).Run(d.Session)
	if err != nil {
		return nil, err
	}
	err = cursor.One(u)
	return
}

// FindUserByID 通过 ID 查找用户
func (d *Dao) FindUserByName(name string) (users []*model.User, err error) {
	cursor, err := r.Table("users").Filter(r.Row.Field("name").Eq(name)).Run(d.Session)
	if err != nil {
		return nil, err
	}
	err = cursor.All(users)
	return
}
