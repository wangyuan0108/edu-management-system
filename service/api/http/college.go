package http

import (
	"context"
	"edu-management-system/db"
	"edu-management-system/schema"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"os"
)

func College() (schema.Status, error) {
	dbBasic := db.Mongo().
		Database(os.Getenv("MONGODB_DB_EDU")).
		Collection(schema.College{}.Collection())
	filter := bson.D{{}}
	cursor, err := dbBasic.Find(context.TODO(), filter)

	var collegeList = make([]map[string]string, 0)
	if err = cursor.All(context.TODO(), &collegeList); err != nil {
		return schema.Status{
			Body:    nil,
			Code:    http.StatusNotAcceptable,
			Message: "请求所有群聊失败",
		}, nil
	}
	return schema.Status{
		Code:    http.StatusOK,
		Message: "请求所有群聊成功",
		Body:    collegeList,
	}, nil
}
