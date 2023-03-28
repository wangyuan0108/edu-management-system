package test

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"testing"
)

func TestMongoDB(t *testing.T) {
	defer func() {
		if err := Mongo().Disconnect(context.TODO()); err != nil {
			t.Fatal(err)
		}
	}()

	// Ping the primary
	if err := Mongo().Ping(context.TODO(), readpref.Primary()); err != nil {
		t.Fatal(err)
	}

	fmt.Println("Successfully connected mongodb and pinged.")
}

func Mongo() *mongo.Client {
	ENV, err := godotenv.Read()
	if err != nil {
		log.Fatal("获取环境变量失败:", err.Error())
	}

	client, collectionErr := mongo.Connect(context.TODO(), options.Client().SetAuth(options.Credential{
		Username: ENV["MONGODB_USERNAME"],
		Password: ENV["MONGODB_PASSWORD"],
	}).ApplyURI(ENV["MONGODB_URL"]))
	if collectionErr != nil {
		log.Println("运行Mongodb服务失败!" + collectionErr.Error())
	}
	return client
}

func TestProductionFind(t *testing.T) {
	ENV, err := godotenv.Read()
	t.Log("MONGODB_URL", ENV["MONGODB_URL"])
	if err != nil {
		log.Fatal("获取环境变量失败:", err.Error())
	}
	client, collectionErr := mongo.Connect(context.TODO(), options.Client().SetAuth(options.Credential{
		Username: ENV["MONGODB_USERNAME"],
		Password: ENV["MONGODB_PASSWORD"],
	}).ApplyURI(ENV["MONGODB_URL"]))
	if collectionErr != nil {
		t.Error(collectionErr)
	}
	find, err := client.
		Database("edu_system").
		Collection("users").
		Find(context.Background(), bson.D{{}})
	var list []bson.M
	if err = find.All(context.Background(), &list); err != nil {
		t.Error(err)
	}
	if err != nil {
		t.Error(err)
	}

	t.Log(list)
}
func TestLocalFind(t *testing.T) {
	ENV, err := godotenv.Read()

	if err != nil {
		log.Fatal("获取环境变量失败:", err.Error())
	}
	client, collectionErr := mongo.Connect(context.TODO(), options.Client().SetAuth(options.Credential{
		Username: ENV["MONGODB_USERNAME"],
		Password: ENV["MONGODB_PASSWORD"],
	}).ApplyURI("mongodb://192.168.0.158:27017"))
	if collectionErr != nil {
		t.Error(collectionErr)
	}
	find, err := client.
		Database("edu_system").
		Collection("users").
		Find(context.Background(), bson.D{{}})
	var list []bson.M
	if err = find.All(context.Background(), &list); err != nil {
		t.Error(err)
	}
	if err != nil {
		t.Error(err)
	}

	t.Log(list)
}
