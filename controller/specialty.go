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
}

type Specialty struct{}

func (Specialty) GET(c *gin.Context) {
	query := c.Query("query")

	if query == "all" {
		condition = bson.D{{}}
		result, err = service.Specialty{}.GetList(condition)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, schema.Status{
				Code:    400,
				Message: "查询异常",
				Body:    nil,
			})
			return
		}
		c.JSON(http.StatusOK, schema.Status{
			Code:    http.StatusOK,
			Message: "获取成功",
			Body:    result,
		})
		return
	} else {
		condition = bson.D{{"name", query}}
		result, err = service.Specialty{}.GetName(condition)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, schema.Status{
				Code:    400,
				Message: "查询异常",
				Body:    nil,
			})
			return
		}
		c.JSON(http.StatusOK, schema.Status{
			Code:    http.StatusOK,
			Message: "获取成功",
			Body:    result,
		})
	}
}

//func (CollegeWithSpecialty) PATCH(c *gin.Context) {
//	var updateData schema.UpdateCollegeWithSpecialty
//	err := c.ShouldBind(&updateData)
//	if err != nil {
//		c.AbortWithStatusJSON(http.StatusOK, schema.Status{
//			Code:    400,
//			Message: "请输入正确的参数",
//			Body:    err.Error(),
//		})
//		return
//	}
//
//	result, patchErr := service.College{}.PATCH(c, updateData)
//	if patchErr != nil {
//		return
//	}
//
//	// 更新专业表下的学院字段所匹配的值为新的值
//	filterSpecialty := bson.D{{"college", updateData.OldCollege}}
//	updateSpecialty := bson.M{"$set": bson.M{"college": updateData.NewCollege}}
//	specialty, specialtyErr := service.UpdateSpecialtyWithCollegeList(filterSpecialty, updateSpecialty)
//	if specialtyErr != nil {
//		c.AbortWithStatusJSON(http.StatusOK, schema.Status{
//			Code:    404,
//			Message: "处理更改专业的请求异常",
//			Body:    err.Error(),
//		})
//		return
//	}
//
//	// 更新专业名称
//	updateSpecialtyName := bson.M{"$set": bson.M{"college": updateData.NewCollege}}
//	specialtyName, SpecialtyNameErr := service.UpdateSpecialtyWithCollegeList(filterSpecialty, updateSpecialtyName)
//	if SpecialtyNameErr != nil {
//		c.AbortWithStatusJSON(http.StatusOK, schema.Status{
//			Code:    404,
//			Message: "处理更改专业的请求异常",
//			Body:    err.Error(),
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, schema.Status{
//		Code:    200,
//		Message: "更新数据成功",
//		Body: gin.H{
//			"specialty":     specialty,
//			"specialtyName": specialtyName,
//			"college":       result,
//		},
//	})
//
//}
