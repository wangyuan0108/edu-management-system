package controller

import (
	"edu-management-system/service/api/oss"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		log.Println("存储文件失败" + err.Error())
	}

	// 上传目录
	dst := "./uploads/"
	if err := c.SaveUploadedFile(file, dst+file.Filename); err != nil {
		log.Println("存储文件失败" + err.Error())
	}

	// 上传文件至minio
	uploadInfo := oss.UploadFile(file.Filename)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": "上传文件至文件存储服务器失败:" + err.Error(),
			"body":    nil,
			"code":    http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "上传文件至文件存储服务器成功",
		"body":    uploadInfo,
		"code":    http.StatusOK,
	})
}
