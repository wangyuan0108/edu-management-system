package http

import (
	"context"
	"edu-management-system/db"
	"edu-management-system/helper"
	"edu-management-system/schema"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
	"time"
)

var token string       // 用于保存JWT令牌
var claim schema.Claim // JWT的标准加密方式
var jwtKey []byte      // 用于签名的私钥

func RoleAuth(person schema.LoginBasic) (schema.Status, error) {
	jwtKey = []byte(os.Getenv("JWT_KEY")) // 读取配置文件获取jwtKey私钥
	issuer := os.Getenv("ISSUER")         // 读取配置文件获取签发者

	//与数据库的用户信息比对 s
	var databasePerson schema.UserBasic // 用户的数据库结构体

	dbName := os.Getenv("MONGODB_DB_EDU") // 读取数据库名
	userColl := schema.UserBasic{}.Collection()
	filter := bson.D{{"account", person.Account}}                           //数据库查询条件
	coll := db.Mongo().Database(dbName).Collection(userColl)                // 配置查询
	findErr := coll.FindOne(context.TODO(), filter).Decode(&databasePerson) // 查询并赋给databasePerson用户的数据库结构体

	// 查询失败情况: 数据库异常
	if findErr != nil {
		if findErr == mongo.ErrNoDocuments {
			return schema.Status{
				Code:    http.StatusBadRequest,
				Message: "查询失败,查询无结果:" + findErr.Error(),
				Body:    nil,
			}, nil
		}
		return schema.Status{
			Code:    http.StatusBadRequest,
			Message: "查询失败,数据库不存在该字段值:" + findErr.Error(),
			Body:    nil,
		}, findErr
	}

	// JWT结构体赋值
	claim = schema.Claim{
		Account: person.Account,      // 用户ID
		Role:    databasePerson.Role, // 用户权限
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1296000 * time.Second).Unix(), // 15天过期
			Issuer:    issuer,                                       // 签发者
		},
	}

	// 生成token
	var generateErr error
	token, generateErr = helper.GenerateToken(jwt.SigningMethodHS256, claim, jwtKey)
	if generateErr != nil || token == "" {
		return schema.Status{
			Code:    http.StatusBadRequest,
			Message: "生成token失败:" + generateErr.Error(),
			Body:    nil,
		}, generateErr
	}

	// 根据前端传递的用户名与用户的数据库结构体的用户名比对, 正确继续下一步-
	if person.Account == databasePerson.Account {
		// 比对前端传递的密码与数据库的密码
		if person.Password == databasePerson.Password {
			// 返回Token
			return schema.Status{
				Code:    http.StatusOK,
				Message: "登录成功",
				Body: struct {
					Token string           `json:"token"`
					Data  schema.UserBasic `json:"data"`
				}{
					Token: token,
					Data:  databasePerson,
				},
			}, nil
		}
		// 密码错误情况
		return schema.Status{
			Code:    http.StatusBadRequest,
			Message: "验证失败,密码错误",
			Body:    nil,
		}, errors.New("验证失败,密码错误")
	}

	// 前端用户名与数据库的用户名不匹配情况
	return schema.Status{
		Code:    http.StatusNotFound,
		Message: "账号输入错误,不存在该账号",
		Body:    nil,
	}, errors.New("账号输入错误,不存在该账号")
}

func ParseJWT(c *gin.Context) {
	token, err := helper.ParseToken(token, jwtKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, schema.Status{
			Code:    http.StatusForbidden,
			Message: "校验JWT失败",
			Body:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, schema.Status{
		Code:    http.StatusOK,
		Message: "校验成功",
		Body:    token,
	})
}
