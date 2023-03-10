package controller

import (
	"edu-management-system/schema"
	api "edu-management-system/service/api/http"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	// 创建用户结构体,用于绑定JSON数据
	var userInfo schema.UserStudentBasic
	flag := true

	file, err := c.FormFile("avatar")
	if file != nil {
		if err != nil {
			flag = false
			log.Println("存储文件失败" + err.Error())
			return
		}
		dst := "./uploads/"
		if err := c.SaveUploadedFile(file, dst+file.Filename); err != nil {
			log.Println("存储文件失败" + err.Error())
			flag = false
		}
	}

	// JSON转为结构体,转换失败抛出异常给客户端与输出异常
	if err := c.ShouldBind(&userInfo); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"data":    "转换JSON失败",
			"message": err.Error(),
			"code":    http.StatusBadRequest,
		})
		log.Println("转换JSON失败")
		return
	}

	var result any
	var registerErr error
	role := c.PostForm("role")
	if !flag {
		_, registerErr := api.Register(userInfo, nil, role)
		if registerErr != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"data":    http.StatusBadRequest,
				"message": "创建用户异常",
				"code":    http.StatusOK,
			})
			log.Println("创建用户异常" + err.Error())
			return
		}
	}
	result, registerErr = api.Register(userInfo, file, role)

	if registerErr != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"data":    http.StatusBadRequest,
			"message": "创建用户异常",
			"code":    http.StatusOK,
		})
		log.Println("创建用户异常" + err.Error())
		return
	}

	// 成功返回客户端创建好的数据条目的ID
	c.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": "注册成功",
		"code":    http.StatusOK,
	})
}
