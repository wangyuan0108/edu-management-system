package controller

import (
	"edu-management-system/schema"
	service "edu-management-system/service/api/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func Specialty(c *gin.Context) {
	query := c.Query("specialty")
	var filter bson.D

	if query == "all" {
		filter = bson.D{{}}
	} else {
		filter = bson.D{{"name", query}}
	}

	result, err := service.GetSpecialtyList(filter)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, schema.Status{
			Code:    404,
			Message: "请求异常",
			Body:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, schema.Status{
		Code:    200,
		Message: "请求专业列表成功",
		Body:    result,
	})
}
