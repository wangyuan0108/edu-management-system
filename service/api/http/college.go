package http

import (
	"context"
	"edu-management-system/db"
	"edu-management-system/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"os"
)

type CollegeInterface interface {
	GetList()
	GetCollegeWithSpecialtyList()
	UpdateCollegeName()
	UpdateCollege()
	AddCollege()
}

type College struct{}

// GetList 获取学院列表
func (College) GetList(filter bson.D) (schema.Status, error) {
	var specialtyList []schema.Specialty
	cursor, err := db.Mongo().
		Database(os.Getenv("MONGODB_DB_EDU")).
		Collection(schema.College{}.Collection()).
		Find(context.Background(), filter)
	if err = cursor.All(context.TODO(), &specialtyList); err != nil {
		log.Println(err)
		return schema.Status{
			Code:    500,
			Message: "数据库异常",
			Body:    err.Error(),
		}, err
	}

	return schema.Status{
		Code:    200,
		Message: "获取学院列表成功",
		Body:    specialtyList,
	}, nil
}

// GetCollegeWithSpecialtyList 学院表关联专业表, 返回每个学院的专业信息
func (College) GetCollegeWithSpecialtyList(pipeline []bson.M) (schema.Status, error) {
	var collegeUnionList []bson.M

	cursor, err := db.Mongo().
		Database(os.Getenv("MONGODB_DB_EDU")).
		Collection(schema.College{}.Collection()).
		Aggregate(context.Background(), pipeline)
	if err != nil {
		return schema.Status{
			Code:    500,
			Message: "数据库处理异常",
			Body:    err.Error(),
		}, err
	}

	if err = cursor.All(context.Background(), &collegeUnionList); err != nil {
		log.Println("数据库处理异常:", err)
		return schema.Status{
			Code:    500,
			Message: "数据库处理异常",
			Body:    err.Error(),
		}, err
	}

	return schema.Status{
		Code:    200,
		Message: "联合查询成功",
		Body:    collegeUnionList,
	}, nil
}

// UpdateCollegeName 更改学院名称
func (College) UpdateCollegeName(OldCollegeName, NewCollegeName string) (schema.Status, error) {
	filter := bson.D{{"name", OldCollegeName}}
	update := bson.M{"$set": bson.M{"name": NewCollegeName}}

	result, err := db.Mongo().
		Database(os.Getenv("MONGODB_DB_EDU")).
		Collection(schema.College{}.Collection()).
		UpdateOne(context.Background(), filter, update)
	if err != nil {
		return schema.Status{}, err
	}

	return schema.Status{
		Code:    200,
		Message: "修改成功",
		Body:    result,
	}, nil
}

// UpdateCollege 更新学院信息
func (College) UpdateCollege(college schema.UpdateCollege) (schema.Status, error) {
	filter := bson.D{{"name", college.OldName}}
	update := bson.M{"$set": bson.M{"name": college.NewName, "description": college.Description}}

	result, err := db.Mongo().
		Database(os.Getenv("MONGODB_DB_EDU")).
		Collection(schema.College{}.Collection()).
		UpdateOne(context.Background(), filter, update)
	if err != nil {
		return schema.Status{}, err
	}

	return schema.Status{
		Code:    200,
		Message: "修改成功",
		Body:    result,
	}, nil
}

// AddCollegeOne 更新学院信息
func (College) AddCollegeOne(college schema.College) (schema.Status, error) {
	var collegeList []schema.College
	insert := bson.D{{"name", college.Name}}
	filter := bson.D{{}}
	findResult, queryErr := db.Mongo().
		Database(os.Getenv("MONGODB_DB_EDU")).
		Collection(schema.College{}.Collection()).
		Find(context.Background(), filter)
	if queryErr != nil {
		return schema.Status{}, queryErr
	}

	if err := findResult.All(context.Background(), &collegeList); err != nil {
		return schema.Status{}, err
	}

	for _, v := range collegeList {
		if v.Name == college.Name {
			return schema.Status{
				Code:    http.StatusConflict,
				Message: "字段的值冲突",
				Body:    nil,
			}, nil
		}
	}

	result, err := db.Mongo().
		Database(os.Getenv("MONGODB_DB_EDU")).
		Collection(schema.College{}.Collection()).
		InsertOne(context.Background(), insert)
	if err != nil {
		return schema.Status{}, err
	}

	return schema.Status{
		Code:    200,
		Message: "添加成功",
		Body:    result,
	}, nil
}

// DeleteCollegeOne 更新学院信息
func (College) DeleteCollegeOne(filter bson.D) (schema.Status, error) {
	result, err := db.Mongo().
		Database(os.Getenv("MONGODB_DB_EDU")).
		Collection(schema.College{}.Collection()).
		DeleteOne(context.Background(), filter)
	if err != nil {
		return schema.Status{}, err
	}
	return schema.Status{
		Code:    200,
		Message: "删除成功",
		Body:    result,
	}, err
}

func (College) GetCollegeOne(filter bson.D) (schema.Status, error) {
	var collegeInfo schema.College
	err := db.Mongo().
		Database(os.Getenv("MONGODB_DB_EDU")).
		Collection(schema.College{}.Collection()).
		FindOne(context.Background(), filter).Decode(&collegeInfo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return schema.Status{
				Code:    404,
				Message: "查询的内容在数据中不存在",
				Body:    nil,
			}, err
		}
		return schema.Status{
			Code:    400,
			Message: "参数异常",
			Body:    nil,
		}, err
	}
	return schema.Status{
		Code:    200,
		Message: "查询成功",
		Body:    collegeInfo,
	}, nil
}
