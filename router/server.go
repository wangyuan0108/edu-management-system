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
	server.Any("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Pong")
	})
	server.POST("/auth", controller.Auth)
	// 路由守卫
	userGroup := server.Group("auth", middleware.AuthRequired)
	{
		// 学院路由
		userGroup.GET("/college", controller.College{}.GET)
		userGroup.PATCH("/college", controller.College{}.PATCH)
		userGroup.POST("/college", controller.College{}.POST)
		userGroup.DELETE("/college", controller.College{}.DELETE)
		userGroup.PUT("/college", controller.College{}.PUT)

		// 上传路由
		userGroup.PUT("/upload", controller.Upload)

		// 专业路由
		userGroup.GET("/specialty", controller.Specialty{}.GET)
		userGroup.PATCH("/specialty", controller.Specialty{}.PATCH)

		// 班级路由
		userGroup.GET("/class", controller.Class{}.GET)

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
