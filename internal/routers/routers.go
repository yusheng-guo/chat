package routers

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	v1 "github.com/yushengguo557/chat/api/v1"
	"github.com/yushengguo557/chat/docs"
	"github.com/yushengguo557/chat/global"
	"github.com/yushengguo557/chat/internal/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	// r.Use(gin.Logger())   // 使用Logger中间件
	r.Use(gin.Recovery()) // 使用Recovery中间件
	// gin.DefaultWriter = global.Logger.Writer()
	r.Use(gin.LoggerWithWriter(global.Logger.Writer())) // 自定义中间件
	docs.SwaggerInfo.BasePath = "http://localhost:8080/"
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// http://127.0.0.1:8080/docs/index.html
	apiv1 := r.Group("/v1") // 路由组
	{
		apiv1.POST("/register", v1.Register)
		apiv1.POST("/login", v1.Login)
		apiv1.POST("/admin", v1.Admin)
		apiv1.POST("/logout", v1.Logout)
		apiv1.PUT("/me", v1.UpdateMyInfo)
		apiv1.GET("/me", v1.GetMyInfo)

		apiv1.GET("/friends", v1.GetMyFriends)
		apiv1.POST("/friend/:id", v1.AddFriend)
		apiv1.DELETE("/friend/:id", v1.DeleteFriend)
		apiv1.PUT("/friend/:id", v1.UpdateFriendInfo)
		apiv1.GET("/friend/:id", v1.GetFriendInfo)

		apiv1.POST("/message", v1.SendMessage)
		apiv1.DELETE("/message/:id", v1.DeleteMessage)
		apiv1.PUT("/message/:id", v1.UpdateMessage)
		apiv1.GET("/message", v1.ReceiveMessage)

		apiv1.GET("/ws", middleware.JWTAuthMiddleware(), v1.HandleWebSocket)

		// apiv1.Static("/image", "./upload")
		apiv1.POST("/upload/image", v1.UploadImage)
	}
	return r
}
