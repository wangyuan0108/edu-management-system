package http

import (
	"edu-management-system/helper"
	"edu-management-system/schema"
	"go.mongodb.org/mongo-driver/bson"
)

type Class struct{}

type ClassInterface interface {
	GetList()
}

func (Class) GetList(query string) (schema.Status, error) {
	filter := bson.D{{"specialty", query}}

	var classList []schema.Class

	//findResult, queryErr := db.Mongo().
	//	Database(os.Getenv("MONGODB_DB_EDU")).
	//	Collection(schema.Class{}.Collection()).
	//	Find(context.Background(), filter)
	//
	//if err := findResult.All(context.Background(), &classList); err != nil {
	//	return schema.Status{
	//		Code:    500,
	//		Message: "服务器处理班级列表失败:" + queryErr.Error(),
	//		Body:    nil,
	//	}, err
	//}
	//
	//if queryErr != nil {
	//	if queryErr == mongo.ErrNoDocuments {
	//		return schema.Status{
	//			Code:    400,
	//			Message: "请求的专业没有对应的班级信息:" + queryErr.Error(),
	//			Body:    nil,
	//		}, queryErr
	//	}
	//	return schema.Status{
	//		Code:    404,
	//		Message: "请求班级列表失败:" + queryErr.Error(),
	//		Body:    nil,
	//	}, queryErr
	//}
	//
	//return schema.Status{
	//	Code:    200,
	//	Message: "请求班级列表成功",
	//	Body:    classList,
	//}, queryErr
	find, err := helper.Mongo{}.Find(schema.Class{}.Collection(), schema.Class{}.Name(), filter, classList)
	if err != nil {
		return schema.Status{}, err
	}

	return find, nil
}
