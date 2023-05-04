package dao

import (
	"fmt"
	"log"
	"testing"

	"github.com/yushengguo557/chat/global"
	"github.com/yushengguo557/chat/internal/model"
)

// TestInsertUser 测试 消息 增删改查
// {0xc000412c00 e80c0552-3027-414c-b208-4031fe730f3d 5547abe4-386e-45c0-af0a-ad9cc9b82e59 hi 0 0}
func TestInsertMessage(t *testing.T) {
	msg := model.NewMessage("e80c0552-3027-414c-b208-4031fe730f3d", "5547abe4-386e-45c0-af0a-ad9cc9b82e59", "test")
	global.InitDB() // 初始化数据库
	d := NewDao(global.Session, global.RedisClient)

	// C
	err := d.InsertMessage(msg)
	if err != nil {
		log.Panic(err)
	}
	log.Println("消息插入")

	// U
	err = d.UpdatemessageByID("c6c3c4c3-d2f6-4910-8fac-6d52150dced7", map[string]any{"content": "hi!😏"})
	if err != nil {
		log.Panic(err)
	}
	log.Println("消息更新")

	// D
	err = d.DeleteMessageByID(msg.ID)
	if err != nil {
		log.Panic(err)
	}
	log.Println("消息删除")
}

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
	// global.InitDB() // 初始化数据库
	// d := NewDao(global.Session, global.RedisClient)
	// err := d.AddOnlineUser("123")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("添加成功！")

	// is := d.IsOnline("123")
	// fmt.Println("是否在线：", is)

	// var u *model.OnlineUser
	// u, err = d.GetOnlineUser("123")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("查询成功！", *u)

	// err = d.RemoveOnlineUser("123")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("删除成功！")
}
