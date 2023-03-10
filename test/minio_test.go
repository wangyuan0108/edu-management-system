package test

import (
	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"testing"
)

var MinioClient *minio.Client

func TestMinioNewClient(t *testing.T) {
	ENV, err := godotenv.Read()
	if err != nil {
		log.Fatal("获取环境变量失败:", err.Error())
	}

	// 初使化minio client对象。
	MinioClient, err = minio.New(ENV["MINIO_URL"], &minio.Options{
		Creds:  credentials.NewStaticV4(ENV["MINIO_ACCESS_KEY"], ENV["MINIO_SECRET_KEY"], ""),
		Secure: false,
	})
	if err != nil {
		t.Fatal("初始化Minio客户端失败" + err.Error())
	}
	t.Log("测试初始化Minio客户端成功", MinioClient)

}
