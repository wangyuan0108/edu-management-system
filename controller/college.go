package controller

import (
	"edu-management-system/schema"
	service "edu-management-system/service/api/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

type CollegeInterface interface {
	GET()
	PATCH()
	PUT()
	POST()
	DELETE()
}

type College struct{}

func (College) GET(c *gin.Context) {
	query := c.Query("query")

	if query == "specialty" {
		pipeline = []bson.M{
			{
				"$lookup": bson.M{
					"from":         "specialty",
					"localField":   "name",
					"foreignField": "college",
					"as":           "college_list",
				},
			},
		}
		result, err = service.College{}.GetCollegeWithSpecialtyList(pipeline)
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
	} else if query == "all" {
		filter := bson.D{{}}
		specialties, err := service.College{}.GetList(filter)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, err)
			return
		}
		c.JSON(http.StatusOK, specialties)
	} else {
		filter := bson.D{{"name", query}}
		specialties, err := service.College{}.GetCollegeOne(filter)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, err)
			return
		}
		c.JSON(http.StatusOK, specialties)
	}
}

func (College) PATCH(c *gin.Context) {
	var updateData schema.SpecialtyWithCollege
	if err := c.ShouldBindWith(&updateData, binding.Form); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, schema.Status{
			Code:    400,
			Message: "参数异常",
			Body:    nil,
		})
		log.Println("绑定参数失败")
		return
	}

	// 更新学院名称
	collegeResult, collegeNameErr := service.College{}.UpdateCollegeName(updateData.OldCollegeName, updateData.NewCollegeName)
	if collegeNameErr != nil {
		c.AbortWithStatusJSON(http.StatusOK, schema.Status{
			Code:    400,
			Message: "请求异常",
			Body:    nil,
		})
		log.Println("绑定参数失败")
		return
	}

	// 更新专业表的学院字段, 对应新修改的学院名称
	specialtyResult, specialtyInfoErr := service.Specialty{}.UpdateSpecialtyWithCollege(updateData.OldCollegeName, updateData.NewCollegeName)
	if specialtyInfoErr != nil {
		c.AbortWithStatusJSON(http.StatusOK, schema.Status{
			Code:    400,
			Message: "更新专业表的学院字段失败",
			Body:    nil,
		})
		log.Println("更新专业表的学院字段失败")
		return
	}

	c.JSON(http.StatusOK, schema.Status{
		Code:    http.StatusOK,
		Message: "获取成功",
		Body: gin.H{
			"collegeResult":   collegeResult,
			"specialtyResult": specialtyResult,
		},
	})
}

// PUT 添加学院信息
func (College) PUT(c *gin.Context) {
	var college schema.UpdateCollege
	if err := c.ShouldBindJSON(&college); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, schema.Status{
			Code:    400,
			Message: "参数异常",
			Body:    nil,
		})
		log.Println("绑定参数失败")
		return
	}

	log.Println("college:", college)
	result, err := service.College{}.UpdateCollege(college)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, schema.Status{
			Code:    400,
			Message: "更新学院表字段值失败",
			Body:    nil,
		})
		log.Println("更新专业表的学院字段失败")
		return
	}

	c.JSON(http.StatusCreated, schema.Status{
		Code:    http.StatusCreated,
		Message: "更新全部数据成功",
		Body:    result,
	})
}

// POST 更新学院信息
func (College) POST(c *gin.Context) {
	var college schema.College
	if err := c.ShouldBindJSON(&college); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, schema.Status{
			Code:    400,
			Message: "参数异常",
			Body:    nil,
		})
		log.Println("绑定参数失败")
		return
	}

	result, err := service.College{}.AddCollegeOne(college)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, schema.Status{
			Code:    400,
			Message: "更新学院表字段值失败",
			Body:    nil,
		})
		log.Println("更新专业表的学院字段失败")
		return
	}

	c.JSON(http.StatusOK, result)
}

// DELETE 删除学院某个文档
func (College) DELETE(c *gin.Context) {
	value := c.Query("delete")
	// 过滤条件
	filter := bson.D{{"name", value}}

	result, err := service.College{}.DeleteCollegeOne(filter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, result)
		return
	}

	c.JSON(http.StatusOK, result)
}
