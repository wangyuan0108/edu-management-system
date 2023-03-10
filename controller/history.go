package controller

import (
	"edu-management-system/db"
	"edu-management-system/schema"
	api "edu-management-system/service/api/http"
	"github.com/gin-gonic/gin"
	"net/http"
)

func History(c *gin.Context) {
	account := c.Query("account")

	api.HistoryMessage(account)

	result, err := db.Redis.LRange(account, 0, -1).Result()
	if err != nil {
		c.JSON(http.StatusOK, schema.Status{
			Code:    http.StatusInternalServerError,
			Message: "获取消息记录失败",
			Body:    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, result)
}
