package http

import (
	"context"
	"edu-management-system/db"
	"edu-management-system/schema"
	"edu-management-system/service/api/oss"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"mime/multipart"
	"os"
)

func Register(userInfo schema.UserStudentBasic, file *multipart.FileHeader, role string) (any, error) {

	log.Printf("userInfo%#v", userInfo)
	// 创建用户
	result, err := db.Mongo().
		Database(os.Getenv("MONGODB_DB_EDU")).
		Collection(schema.UserStudentBasic{}.Collection()).
		InsertOne(context.TODO(), userInfo)
	log.Println("result:", result)
	// 创建用户时的异常处理
	if err != nil {
		return nil, err
	}

	if file != nil {
		go oss.UploadAvatar(file.Filename, result.InsertedID.(primitive.ObjectID), role)
		return result.InsertedID, err
	}
	return result.InsertedID, nil
}
