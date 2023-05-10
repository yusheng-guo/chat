// 说明：朋友管理
package service

import (
	"fmt"
)

// AddFriendByID 通过用户id 添加好友
func (s *Service) AddFriendByID(myid string, friendid string) error {
	// 从 rethinkdb 中查找 me 和 friend
	me, err := s.dao.FindUserByID(myid)
	if err != nil {
		return err
	}
	friend, err := s.dao.FindUserByID(friendid)
	if err != nil {
		return err
	}
	// 更新 rethinkdb 数据库
	me.Friends[friendid] = friend.Name
	friend.Friends[myid] = me.Name
	err = s.dao.UpdateUserByID(myid, map[string]any{"friends": me.Friends})
	if err != nil {
		return fmt.Errorf("update user by id in `AddFriendByID` Function")
	}
	err = s.dao.UpdateUserByID(friendid, map[string]any{"friends": friend.Friends})
	if err != nil {
		return fmt.Errorf("update user by id in `AddFriendByID` Function")
	}
	return nil
}
