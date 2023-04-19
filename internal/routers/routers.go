package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/yushengguo557/chat/api"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())   // 使用Logger中间件
	r.Use(gin.Recovery()) // 使用Recovery中间件
	a := r.Group("/a")
	{
		a.POST("/register", api.Register)
		a.POST("/login", api.Login)
		a.POST("/admin", api.Admin)
		a.POST("/logout", api.Logout)

		a.PUT("/me/:id/", api.UpdateMyInfo)
		a.GET("/me/:id/", api.GetMyInfo)

		a.GET("/me/:id/friends", api.GetMyFriends)
		a.POST("/friend/:id", api.AddFriend)
		a.DELETE("/friend/:id", api.DeleteFriend)
		a.PUT("/friend/:id", api.UpdateFriendInfo)
		a.GET("/friend/:id", api.GetFriendInfo)

		a.POST("/message", api.SendMessage)
		a.DELETE("/message/:id", api.DeleteMessage)
		a.PUT("/message/:id", api.UpdateMessage)
		a.GET("/message", api.ReceiveMessage)
	}
	return r
}
