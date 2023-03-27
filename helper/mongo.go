package helper

import (
	"context"
	"edu-management-system/db"
	"edu-management-system/schema"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

type methods interface {
	Find(table string, dbname string, filter any, result any) (schema.Status, error)
	FindOne(table string, dbname string, filter any, result any) (schema.Status, error)
	InsertOne(table string, dbname string, data any) (schema.Status, error)
	InsertMany(table string, dbname string, data []any) (schema.Status, error)
	UpdateOne(table string, dbname string, filter any, update any) (schema.Status, error)
	UpdateMany(table string, dbname string, filter any, update any) (schema.Status, error)
	DeleteOne(table string, dbname string, filter any) (schema.Status, error)
	DeleteMany(table string, dbname string, filter any) (schema.Status, error)
}

type Mongo struct {
	methods
}

// Find
/* @description 全部查询方法
 * @since 27/03/2023 11:09 am
 * @param table 数据库名称
 * @param dbname 数据库集合
 * @param filter 查询条件
 * @param result 查询结果
 * @return { schema.Status } 返回状态
 * @Error { error } 错误信息
 *  */
func (Mongo) Find(table string, dbname string, filter any, result any) (schema.Status, error) {
	// 数据库配置
	findResult, queryErr := db.Mongo().
		Database(os.Getenv("MONGODB_DB_EDU")).
		Collection(table).
		Find(context.Background(), filter)

	// 对数据库返回的数组进行处理
	if err := findResult.All(context.Background(), &result); err != nil {
		return schema.Status{
			Code:    500,
			Message: "服务器处理%s失败:" + queryErr.Error(),
			Body:    nil,
		}, err
	}

	// 关闭Find查询连接
	defer func(findResult *mongo.Cursor, ctx context.Context) {
		err := findResult.Close(ctx)
		if err != nil {
			log.Printf("关闭查询%s失败, 错误原因:%s", dbname, err.Error())
		}
	}(findResult, context.Background())

	// 判断查询结果
	if queryErr != nil {
		// 判断是否没有对应的信息
		if queryErr == mongo.ErrNoDocuments {
			log.Printf("请求%s失败,没有对应的信息, 错误原因:%s", dbname, queryErr.Error())
			return schema.Status{
				Code:    400,
				Message: fmt.Sprintf("请求%s失败,没有对应的信息, 错误原因:%s", dbname, queryErr.Error()),
				Body:    nil,
			}, queryErr
		}
		// 其他错误
		return schema.Status{
			Code:    404,
			Message: fmt.Sprintf("请求%s失败,错误原因:%s", dbname, queryErr.Error()),
			Body:    nil,
		}, queryErr
	}

	// 返回查询结果
	return schema.Status{
		Code:    200,
		Message: fmt.Sprintf("请求%s成功", dbname),
		Body:    result,
	}, queryErr
}

// FindOne
/* @description 查询单条记录
 * @since 27/03/2023 11:17 am
 * @param table 数据库名称
 * @param dbname 数据库集合
 * @param filter 查询条件
 * @param result 查询结果
 * @return { schema.Status } 返回状态
 * @Error { error } 错误信息
 *  */
func (Mongo) FindOne(table string, dbname string, filter any, result any) (schema.Status, error) {
	findResult := db.Mongo().
		Database(os.Getenv("MONGODB_DB_EDU")).
		Collection(table).
		FindOne(context.Background(), filter)

	err := findResult.Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("请求%s失败,没有对应的信息, 错误原因:%s", dbname, err.Error())
			return schema.Status{
				Code:    404,
				Message: fmt.Sprintf("请求%s失败,没有对应的信息, 错误原因:%s", dbname, err.Error()),
				Body:    nil,
			}, err
		}
		return schema.Status{
			Code:    500,
			Message: fmt.Sprintf("请求%s失败,错误原因:%s", dbname, err.Error()),
			Body:    nil,
		}, err
	}

	return schema.Status{
		Code:    200,
		Message: fmt.Sprintf("请求%s成功", dbname),
		Body:    result,
	}, nil
}

// InsertOne
/* @description 插入单条记录
 * @since 27/03/2023 11:17 am
 * @param table 数据库名称
 * @param dbname 数据库集合
 * @param data 需要插入的数据
 * @return { schema.Status } 返回状态
 * @Error { error } 错误信息
 *  */
func (Mongo) InsertOne(table string, dbname string, data any) (schema.Status, error) {
	insertResult, err := db.Mongo().
		Database(os.Getenv("MONGODB_DB_EDU")).
		Collection(table).
		InsertOne(context.Background(), data)

	if err != nil {
		return schema.Status{
			Code:    500,
			Message: fmt.Sprintf("插入%s失败,错误原因:%s", dbname, err.Error()),
			Body:    nil,
		}, err
	}

	return schema.Status{
		Code:    200,
		Message: fmt.Sprintf("插入%s成功", dbname),
		Body:    insertResult.InsertedID,
	}, nil
}

// InsertMany
/* @description 插入多条记录
 * @since 27/03/2023 11:18 am
 * @param table 数据库名称
 * @param dbname 数据库集合
 * @param data 需要插入的数据
 * @return { schema.Status } 返回状态
 * @Error { error } 错误信息
 *  */
func (Mongo) InsertMany(table string, dbname string, data []any) (schema.Status, error) {
	insertResult, err := db.Mongo().
		Database(os.Getenv("MONGODB_DB_EDU")).
		Collection(table).
		InsertMany(context.Background(), data)

	if err != nil {
		return schema.Status{
			Code:    500,
			Message: fmt.Sprintf("插入%s失败,错误原因:%s", dbname, err.Error()),
			Body:    nil,
		}, err
	}

	return schema.Status{
		Code:    200,
		Message: fmt.Sprintf("插入%s成功", dbname),
		Body:    insertResult.InsertedIDs,
	}, nil
}

// UpdateOne
/* @description 更新单条记录
 * @since 27/03/2023 11:19 am
 * @param table 数据库名称
 * @param dbname 数据库集合
 * @param filter 查询条件
 * @param update 更新内容
 * @return { schema.Status } 返回状态
 * @Error { error } 错误信息
 *  */
func (Mongo) UpdateOne(table string, dbname string, filter any, update any) (schema.Status, error) {
	updateResult, err := db.Mongo().
		Database(os.Getenv("MONGODB_DB_EDU")).
		Collection(table).
		UpdateOne(context.Background(), filter, update)

	if err != nil {
		return schema.Status{
			Code:    500,
			Message: fmt.Sprintf("更新%s失败,错误原因:%s", dbname, err.Error()),
			Body:    nil,
		}, err
	}

	return schema.Status{
		Code:    200,
		Message: fmt.Sprintf("更新%s成功", dbname),
		Body:    updateResult.ModifiedCount,
	}, nil
}

// UpdateMany
/* @description 更新多条记录
 * @since 27/03/2023 11:21 am
 * @param table 数据库名称
 * @param dbname 数据库集合
 * @param filter 查询条件
 * @param update 更新内容
 * @return { schema.Status } 返回状态
 * @Error { error } 错误信息
 *  */
func (Mongo) UpdateMany(table string, dbname string, filter any, update any) (schema.Status, error) {
	updateResult, err := db.Mongo().
		Database(os.Getenv("MONGODB_DB_EDU")).
		Collection(table).
		UpdateMany(context.Background(), filter, update)

	if err != nil {
		return schema.Status{
			Code:    500,
			Message: fmt.Sprintf("更新%s失败,错误原因:%s", dbname, err.Error()),
			Body:    nil,
		}, err
	}

	return schema.Status{
		Code:    200,
		Message: fmt.Sprintf("更新%s成功", dbname),
		Body:    updateResult.ModifiedCount,
	}, nil
}

// DeleteOne
/* @description 删除单条记录
 * @since 27/03/2023 11:22 am
 * @param table 数据库名称
 * @param dbname 数据库集合
 * @param filter 查询条件
 * @return { schema.Status } 返回状态
 * @Error { error } 错误信息
 *  */
func (Mongo) DeleteOne(table string, dbname string, filter any) (schema.Status, error) {
	deleteResult, err := db.Mongo().
		Database(os.Getenv("MONGODB_DB_EDU")).
		Collection(table).
		DeleteOne(context.Background(), filter)

	if err != nil {
		return schema.Status{
			Code:    500,
			Message: fmt.Sprintf("删除%s失败,错误原因:%s", dbname, err.Error()),
			Body:    nil,
		}, err
	}

	return schema.Status{
		Code:    200,
		Message: fmt.Sprintf("删除%s成功", dbname),
		Body:    deleteResult,
	}, nil
}

// DeleteMany
/* @description 删除多条记录
 * @since 27/03/2023 11:24 am
 * @param table 数据库名称
 * @param dbname 数据库集合
 * @param filter 查询条件
 * @return { schema.Status } 返回状态
 * @Error { error } 错误信息
 *  */
func (Mongo) DeleteMany(table string, dbname string, filter any) (schema.Status, error) {
	deleteResult, err := db.Mongo().
		Database(os.Getenv("MONGODB_DB_EDU")).
		Collection(table).
		DeleteMany(context.Background(), filter)

	if err != nil {
		return schema.Status{
			Code:    500,
			Message: fmt.Sprintf("删除%s失败,错误原因:%s", dbname, err.Error()),
			Body:    nil,
		}, err
	}

	return schema.Status{
		Code:    200,
		Message: fmt.Sprintf("删除%s成功", dbname),
		Body:    deleteResult,
	}, nil
}
