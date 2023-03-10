package controller

import (
	service "edu-management-system/service/api/http"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RoomList(c *gin.Context) {
	query := c.Query("number")

	result, err := service.RoomList(query)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, result)
}
