package router

import (
	"edu-management-system/controller"
	"edu-management-system/middleware"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Server() {
	server := gin.Default()
	server.Use(middleware.Cors()) // 跨域处理
	server.PUT("/register", controller.Register)
	server.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Pong")
	})
	server.POST("/auth", controller.Auth)
	// 路由守卫
	userGroup := server.Group("auth", middleware.AuthRequired)
	{
		//userGroup.GET("/", middleware.AuthRequired, api.ParseJWT) // 校验用户的JWT是否正确, 如果校验失败则进入不了用户路由

		//server.POST("/login", controller.Login)
		userGroup.GET("/college", controller.College)
		userGroup.PUT("/upload", controller.Upload)
		userGroup.GET("/specialty", controller.Specialty)

		// 聊天路由
		chatGroup := server.Group("ws")
		{
			chatGroup.GET("/chat", controller.IM)
			chatGroup.GET("/chat/roomList", controller.RoomList)
			chatGroup.GET("/chat/historyMessage", controller.History)
		}
	}

	if err := server.Run("0.0.0.0:4000"); err != nil {
		log.Fatalln("运行gin服务失败,请检查端口是否被占用", err.Error())
	}
}
