package dao

import (
	"fmt"
	"log"
	"testing"

	"github.com/yushengguo557/chat/global"
	"github.com/yushengguo557/chat/internal/model"
)

// TestInsertUser 测试 rethinkdb 数据库
func TestInsertUser(t *testing.T) {
	u := model.NewUser()
	u.Email = "test@gmail.com"
	u.Password = "123"
	global.InitDB() // 初始化数据库
	d := NewDao(global.Session, global.RedisClient)
	err := d.InsertUser(u)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("用户插入成功！")
	_, err = d.FindUserByID(u.ID)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("用户查找成功！")
	err = d.DeleteUserByID(u.ID)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("用户删除成功！")
}

// TestAddOnlineUser 测试 redis 数据库
func TestAddOnlineUser(t *testing.T) {
	global.InitDB() // 初始化数据库
	d := NewDao(global.Session, global.RedisClient)
	err := d.AddOnlineUser("123")
	if err != nil {
		panic(err)
	}
	fmt.Println("添加成功！")

	is := d.IsOnline("123")
	fmt.Println("是否在线：", is)

	var u *model.OnlineUser
	u, err = d.GetOnlineUser("123")
	if err != nil {
		panic(err)
	}
	fmt.Println("查询成功！", *u)

	err = d.RemoveOnlineUser("123")
	if err != nil {
		panic(err)
	}
	fmt.Println("删除成功！")
}
