package oss

import (
	"context"
	"edu-management-system/helper"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"os"
	"path/filepath"
)

// UploadFile 上传文件
/* @description 上传文件至Minio服务器
 * @since 2023/2/205:50
 * @param args[0] 文件名
 * @param args[1] 存储桶名
 * @param args[2] 文件路径
 * @return 文件存储信息
 *  */
func UploadFile(args ...string) minio.UploadInfo {
	var filename, bucketName, filePath string
	if len(args) == 1 {
		uploadDir, _ := helper.GetPath("/uploads/", args[0]) // 待上传的文件路径
		uploadDirToLinux := filepath.FromSlash(uploadDir)    // 转成Minio的服务器路径, 本机部署在Linux,通过FromSlash替换成/路径
		filename = args[0]
		bucketName = "chat"
		filePath = uploadDirToLinux
	} else if len(args) == 3 {
		filename = args[0]
		bucketName = args[1]
		filePath = args[2]
	} else {
		log.Println("请传递正确的参数")
	}

	// 初使化minio client对象。
	MinioClient, err := minio.New(os.Getenv("MINIO_URL"), &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("MINIO_ACCESS_KEY"), os.Getenv("MINIO_SECRET_KEY"), ""),
		Secure: false,
	})
	if err != nil {
		log.Println("初始化Minio客户端失败" + err.Error())
	}

	// 上传文件至Minio服务器
	uploadInfo, uploadErr := MinioClient.FPutObject(
		context.TODO(),
		bucketName,
		filename,
		filePath,
		minio.PutObjectOptions{
			ContentType: "application/octet-stream",
		})
	if uploadErr != nil {
		log.Println("上传文件失败:", err)
	}
	return uploadInfo
}
