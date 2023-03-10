package controller

import (
	service "edu-management-system/service/api/http"
	"github.com/gin-gonic/gin"
	"net/http"
)

func College(c *gin.Context) {
	result, err := service.College()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, result)
	}
	c.JSON(http.StatusOK, result)
}
