package dao

import (
	"fmt"
	"log"
	"testing"

	"github.com/yushengguo557/chat/global"
	"github.com/yushengguo557/chat/internal/model"
)

func TestInsertUser(t *testing.T) {
	u := model.NewUser()
	u.Email = "test"
	u.Password = "123"
	global.InitDB()
	d := New(global.Session)
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
