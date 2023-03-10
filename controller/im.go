package controller

import (
	"edu-management-system/schema"
	ws "edu-management-system/service/api/websocket"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

// 处理WebSocket跨域
var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var webSocketConn = make(map[string]*websocket.Conn)

func IM(c *gin.Context) {
	// 升级为WebSocket协议 s
	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, schema.Status{
			Code:    http.StatusInternalServerError,
			Message: "系统异常：" + err.Error(),
			Body:    nil,
		})
		return
	}

	ws.Chat(c, conn, webSocketConn)
}
