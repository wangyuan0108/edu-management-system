package controller

import (
	"edu-management-system/schema"
	service "edu-management-system/service/api/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

var (
	condition bson.D
	pipeline  []bson.M
	result    any
	err       error
)

type SpecialtyInterface interface {
	GET()
	PATCH()
}

type Specialty struct{}

func (Specialty) GET(c *gin.Context) {
	query := c.Query("query")

	if query == "all" {
		condition = bson.D{{}}
		result, err = service.Specialty{}.GetList(condition)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, err)
			return
		}
		c.JSON(http.StatusOK, result)
	} else {
		condition = bson.D{{"name", query}}
		result, err = service.Specialty{}.GetName(condition)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, err)
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func (Specialty) PATCH(c *gin.Context) {
	var updateData schema.UpdateSpecialty
	err := c.ShouldBindJSON(&updateData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, schema.Status{
			Code:    400,
			Message: "请输入正确的参数",
			Body:    err.Error(),
		})
		return
	}

	filter := bson.D{{"name", updateData.OldSpecialtyName}}
	update := bson.M{"$set": bson.M{"name": updateData.NewSpecialtyName, "description": updateData.NewSpecialtyDescription}}
	result, patchErr := service.Specialty{}.UpdateSpecialty(filter, update)
	if patchErr != nil {
		c.AbortWithStatusJSON(http.StatusOK, patchErr)
		return
	}
	c.JSON(http.StatusOK, result)
}
