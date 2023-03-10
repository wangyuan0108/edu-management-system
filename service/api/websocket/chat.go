package websocket

import (
	"edu-management-system/schema"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

func Chat(c *gin.Context, conn *websocket.Conn, webSocketConn map[string]*websocket.Conn) {
	for {
		// 创建消息JSON结构体, 保存消息与额外的信息
		message := new(schema.MessageBasic)
		// ReadJSON
		// message: 需要读取的消息对象(一条消息一般包含多个属性用于其他用途)
		err := conn.ReadJSON(message)
		if err != nil {
			log.Println("read:", err)
			break
		}

		// 消息结构体
		msg := schema.MessageBasic{
			UserIdentity: message.UserIdentity,
			RoomIdentity: message.RoomIdentity,
			Data:         message.Data,
			CreatedTime:  message.CreatedTime,
			UpdatedTime:  message.UpdatedTime,
		}

		go SaveMessage(c, message)

		// 根据连接绑定用户id
		webSocketConn[msg.UserIdentity] = conn

		for _, conn := range webSocketConn {
			// WriteMessage
			// 1 消息类型: websocket.TextMessage文本
			// 2 传输类型: []byte二进制
			err = conn.WriteJSON(msg)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}
}
