package websocket

import (
	"edu-management-system/db"
	"edu-management-system/schema"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"log"
	"net/http"
)

func SaveMessage(c *gin.Context, msg *schema.MessageBasic) {
	// 转换JSON
	jsonMessage, err := json.Marshal(&msg)
	if err != nil {
		log.Println("转换JSON失败:" + err.Error())
	}

	db.RedisDBInit()
	defer func(client *redis.Client) {
		err := client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db.Redis)

	// 存储消息, 存储异常则返回消息给客户端,继续执行
	if err := db.Redis.RPush(msg.UserIdentity, jsonMessage).Err(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"body":    nil,
			"message": "存储消息失败",
			"code":    http.StatusBadRequest,
		})
		log.Println("存储消息失败:" + err.Error())
	}
}
