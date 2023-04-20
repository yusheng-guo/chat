package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/yushengguo557/chat/api"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())   // 使用Logger中间件
	r.Use(gin.Recovery()) // 使用Recovery中间件
	{
		r.POST("/register", api.Register)
		r.POST("/login", api.Login)
		r.POST("/admin", api.Admin)
		r.POST("/logout", api.Logout)
	}
	{
		r.PUT("/me/:id/", api.UpdateMyInfo)
		r.GET("/me/:id/", api.GetMyInfo)
	}
	{
		r.GET("/me/:id/friends", api.GetMyFriends)
		r.POST("/friend/:id", api.AddFriend)
		r.DELETE("/friend/:id", api.DeleteFriend)
		r.PUT("/friend/:id", api.UpdateFriendInfo)
		r.GET("/friend/:id", api.GetFriendInfo)
	}
	{
		r.POST("/message", api.SendMessage)
		r.DELETE("/message/:id", api.DeleteMessage)
		r.PUT("/message/:id", api.UpdateMessage)
		r.GET("/message", api.ReceiveMessage)
	}
	return r
}
