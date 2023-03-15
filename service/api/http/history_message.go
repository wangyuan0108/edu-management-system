package http

import (
	"edu-management-system/db"
	"edu-management-system/schema"
	"log"
	"net/http"
)

func HistoryMessage(account string) schema.Status {
	result, err := db.Redis.LRange(account, 0, -1).Result()
	if err != nil {
		log.Println("获取消息列表失败" + err.Error())
		return schema.Status{
			Code:    http.StatusBadRequest,
			Body:    result,
			Message: "获取消息列表失败",
		}
	}
	return schema.Status{
		Code:    http.StatusOK,
		Body:    result,
		Message: "获取消息列表成功",
	}
}
