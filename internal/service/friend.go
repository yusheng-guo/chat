// 说明：朋友管理
package service

import (
	"errors"
	"fmt"

	"github.com/yushengguo557/chat/internal/model"
)

// AddFriendByID 通过用户id 添加好友
func (s *Service) AddFriendByID(myid, friendid string) error {
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
	err = s.dao.UpdateUserByID(myid, &map[string]any{"friends": me.Friends})
	if err != nil {
		return fmt.Errorf("update user by id in `AddFriendByID` Function")
	}
	err = s.dao.UpdateUserByID(friendid, &map[string]any{"friends": friend.Friends})
	if err != nil {
		return fmt.Errorf("update user by id in `AddFriendByID` Function")
	}
	return nil
}

// DeleteFriendByID 通过 好友 ID 删除好友
func (s *Service) DeleteFriendByID(myid, friendid string) error {
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
	delete(me.Friends, friendid)
	delete(friend.Friends, myid)
	err = s.dao.UpdateUserByID(myid, &map[string]any{"friends": me.Friends})
	if err != nil {
		return fmt.Errorf("update user by id in `AddFriendByID` Function")
	}
	err = s.dao.UpdateUserByID(friendid, &map[string]any{"friends": friend.Friends})
	if err != nil {
		return fmt.Errorf("update user by id in `AddFriendByID` Function")
	}
	return nil
}

// GetFriendInfoByID 通过 id 获取好友信息
func (s *Service) ModifyFriendNoteByID(myid, friendid, note string) error {
	// 判断是否为好友关系
	u, err := s.dao.FindUserByID(myid)
	if err != nil {
		return err
	}
	_, ok := u.Friends[friendid]
	if !ok {
		return errors.New("you are not friends")
	}
	// 修改备注
	u.Friends[friendid] = note
	if err != nil {
		return err
	}
	return nil
}

// GetFriendInfoByID 通过 id 获取好友信息
func (s *Service) GetFriendInfoByID(myid, friendid string) (u *model.User, err error) {
	// 判断是否为好友关系
	u, err = s.dao.FindUserByID(myid)
	if err != nil {
		return nil, err
	}
	_, ok := u.Friends[friendid]
	if !ok {
		return nil, errors.New("you are not friends")
	}
	// 通过 朋友id 查找朋友
	u, err = s.dao.FindUserByID(friendid)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// GetMyFriends 获取我的所有好友
func (s *Service) GetMyFriends(myid string) ([]string, error) {
	// 从 rethinkdb 中查找 User
	me, err := s.dao.FindUserByID(myid)
	if err != nil {
		return nil, err
	}

	// 获取 所有 好友id
	count := len(me.Friends) // 好友数量
	friends := make([]string, count)
	i := 0
	for id := range me.Friends {
		friends[i] = id
		i++
	}
	return friends, nil
}
