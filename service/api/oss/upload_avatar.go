package oss

import (
	"context"
	"edu-management-system/db"
	"edu-management-system/schema"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"os"
	"path/filepath"
)

// UploadAvatar
/* @description
 * @since 2023/2/2115:50
 * @param args[0] 文件名
 * @param args[1] Mongodb对象ID
 * @param args[2] 角色
 * @return
 *  */
func UploadAvatar(args ...any) {
	go UploadFile(args[0].(string))

	// 协议
	protocol := "http://"
	// 文件存储服务器地址+port
	host := os.Getenv("MINIO_URL")
	// 存储桶名
	bucket := os.Getenv("MINIO_BUCKET")
	// 转为Linux文件路径
	filepathToLinux := filepath.ToSlash(filepath.Join(host, bucket, args[0].(string)))
	// 完整URI
	uri := protocol + filepathToLinux

	dbname := "student"
	switch args[2].(string) {
	case "student":
		dbname = schema.UserStudentBasic{}.Collection()
		break
	case "teacher":
		dbname = schema.UserStudentBasic{}.Collection()
		break
	case "admin":
		dbname = schema.UserAdminBasic{}.Collection()
		break
	default:
		log.Println("非法传入的参数")
	}

	// 对传入的ObjectID文档进行修改, 添加传入的头像URI
	filter := bson.D{{"_id", args[1]}}
	update := bson.D{{"$set", bson.D{{"avatar", uri}}}}
	result, err := db.Mongo().Database(os.Getenv("MONGODB_DB_EDU")).
		Collection(dbname).
		UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
	}
	log.Println(result)
}
