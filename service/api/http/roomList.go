package http

import (
	"context"
	"edu-management-system/db"
	"edu-management-system/schema"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"strconv"
)

func RoomList(query string) (any, error) {
	dbBasic := db.Mongo().
		Database(os.Getenv("MONGODB_DB_IM")).
		Collection(schema.RoomBasic{}.
			Collection())
	if query == "all" { // 查询全部群聊
		// 查询全部
		filter := bson.D{{}}
		// 只返回name字段
		opts := options.Find().SetProjection(bson.D{
			{"info", 0},
			{"user_identity", 0},
			{"created_time", 0},
		})
		cursor, err := dbBasic.Find(context.TODO(), filter, opts)
		var roomList = make([]map[string]any, 0)
		if err = cursor.All(context.TODO(), &roomList); err != nil {
			return schema.Status{
				Body:    nil,
				Code:    http.StatusNotAcceptable,
				Message: "请求所有群聊失败",
			}, nil
		}
		return schema.Status{
			Code:    http.StatusOK,
			Message: "请求所有群聊成功",
			Body:    roomList,
		}, nil
	} else if _, err := strconv.Atoi(query); err != nil { // 如果转为int类型成功,那么为群号直接查询
		// 根据传入的number查询
		var roomInfo schema.RoomBasic
		filter := bson.D{{"number", query}}
		err = dbBasic.FindOne(context.TODO(), filter).Decode(&roomInfo)
		log.Println("roomInfo:", roomInfo)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return schema.Status{
					Body:    nil,
					Code:    http.StatusNotAcceptable,
					Message: "查询的群号无结果",
				}, nil
			}
			log.Println(err)
			return nil, err
		}

		return schema.Status{
			Body:    roomInfo,
			Code:    http.StatusOK,
			Message: fmt.Sprintf("获取number为%v的结果成功", query),
		}, nil
	}
	return schema.Status{
		Body:    nil,
		Code:    http.StatusNotFound,
		Message: "异常/或未携带值的请求参数",
	}, nil
}
