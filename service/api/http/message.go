package http

import (
	"context"
	"edu-management-system/db"
	"edu-management-system/schema"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
)

func MessageList(c *gin.Context) {
	opt := options.Find().SetSort(bson.D{{"created_at", 1}})
	cursor, err := db.Mongo().
		Database(os.Getenv("MONGODB_DB_IM")).
		Collection(schema.MessageBasic{}.Collection()).
		Find(context.TODO(), bson.D{{}}, opt)
	if err != nil {
		log.Println(err)
		return
	}

	var message []schema.MessageBasic

	if err = cursor.All(context.TODO(), &message); err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "获取成功",
		"body":    message,
	})
}
