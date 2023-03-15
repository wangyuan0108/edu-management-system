package http

import (
	"context"
	"edu-management-system/db"
	"edu-management-system/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

type SpecialtyInterface interface {
	GetList()
	GetName()
	UpdateSpecialtyWithCollege()
}

type Specialty struct{}

// GetList 获取专业列表
func (Specialty) GetList(filter any) ([]schema.Specialty, error) {
	var specialtyList []schema.Specialty
	cursor, err := db.Mongo().
		Database(os.Getenv("MONGODB_DB_EDU")).
		Collection(schema.Specialty{}.Collection()).
		Find(context.Background(), filter)
	if err = cursor.All(context.TODO(), &specialtyList); err != nil {
		log.Println(err)
		return nil, err
	}

	return specialtyList, nil
}

// GetName 获取单个专业的信息
func (Specialty) GetName(filter bson.D) (schema.Specialty, error) {
	var specialtyName schema.Specialty
	err := db.Mongo().
		Database(os.Getenv("MONGODB_DB_EDU")).
		Collection(schema.Specialty{}.Collection()).
		FindOne(context.Background(), filter).Decode(&specialtyName)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("未查询到结果")
			return schema.Specialty{}, nil
		}
		return schema.Specialty{}, err
	}
	return specialtyName, nil
}

// UpdateSpecialtyWithCollege 更新学院字段
func (Specialty) UpdateSpecialtyWithCollege(OldCollege, NewCollege string) (schema.Status, error) {
	log.Println(OldCollege)
	log.Println(NewCollege)
	filter := bson.D{{"college", OldCollege}}
	update := bson.M{"$set": bson.M{"college": NewCollege}}

	result, err := db.Mongo().
		Database(os.Getenv("MONGODB_DB_EDU")).
		Collection(schema.Specialty{}.Collection()).
		UpdateMany(context.Background(), filter, update)
	if err != nil {
		return schema.Status{}, err
	}

	return schema.Status{
		Code:    200,
		Message: "修改成功",
		Body:    result,
	}, nil
}
