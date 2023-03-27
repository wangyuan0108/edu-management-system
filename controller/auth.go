package controller

import (
	"edu-management-system/schema"
	service "edu-management-system/service/api/http"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginBasic struct {
	Account  string `bson:"account" json:"account" binding:"required"`
	Password string `bson:"password" json:"password" binding:"required"`
}

func Auth(c *gin.Context) {
	var person schema.LoginBasic // 定义用户结构体并绑定
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusOK, schema.Status{
			Code:    http.StatusBadRequest,
			Message: "JSON绑定失败,请检查你的JSON结构:" + err.Error(),
			Body:    nil,
		})
		return
	}

	result, err := service.RoleAuth(person)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, err)
		return
	}
	c.JSON(http.StatusOK, result)
}
