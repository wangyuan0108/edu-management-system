package controller

import (
	service "edu-management-system/service/api/http"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ClassInterface interface {
	GET(*gin.Context)
	PATCH()
	PUT()
	POST()
	DELETE()
}

type Class struct {
	ClassInterface
}

func (Class) GET(c *gin.Context) {
	query := c.Query("query")

	result, err = service.Class{}.GetList(query)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
