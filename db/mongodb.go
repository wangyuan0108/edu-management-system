package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func Mongo() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().SetAuth(options.Credential{
		Username: os.Getenv("MONGODB_USERNAME"),
		Password: os.Getenv("MONGODB_PASSWORD"),
	}).ApplyURI(os.Getenv("MONGODB_URL")))
	if err != nil {
		log.Println("运行Mongodb服务失败!" + err.Error())
	}
	return client
}
