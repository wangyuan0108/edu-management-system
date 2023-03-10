package schema

import "github.com/dgrijalva/jwt-go"

type LoginBasic struct {
	Account  string `bson:"account" json:"account"`
	Password string `bson:"password" json:"password"`
}
type UserBasic struct {
	Account     string `json:"account" bson:"account" form:"account"  json:"account"`
	Avatar      string `json:"avatar" bson:"avatar" form:"avatar"`
	Password    string `json:"password" bson:"password" form:"password"  json:"password"`
	Username    string `json:"username" bson:"username" form:"username"  json:"username"`
	Role        string `json:"role" bson:"role" form:"role" json:"role"`
	CreatedTime int64  `json:"createdTime" form:"createdTime" bson:"created_time"`
	UpdatedTime int64  `json:"updatedTime" form:"updatedTime" bson:"updated_time"`
}

type UserStudentBasic struct {
	Account        string `bson:"account" form:"account"  json:"account"`
	Password       string `bson:"password" form:"password"  json:"password"`
	Username       string `bson:"username" form:"username"  json:"username"`
	Avatar         string `bson:"avatar" form:"updatedAt" json:"avatar"`
	Role           string `bson:"role" form:"role"`
	Gender         string `bson:"gender" form:"gender"`
	College        string `bson:"college" form:"college"`       // 学院
	Profession     string `bson:"profession" form:"profession"` // 专业
	Class          string `bson:"class" form:"class"`
	EnrollmentDate int64  `bson:"enrollment_date" form:"enrollmentDate"` // 入学日期
	Address        string `bson:"address" form:"address" binding:"required,min=10"`
	Birthdate      int64  `bson:"birthdate" form:"birthdate"`            // 出生日期
	NativePlace    string `bson:"native_place" form:"nativePlace"`       // 籍贯
	PoliticsStatus string `bson:"politics_status" form:"politicsStatus"` // 政治身份
	CreatedTime    int64  `form:"createdTime" bson:"created_time"`
	UpdatedTime    int64  `form:"updatedTime" bson:"updated_time"`
}

type UserTeacherBasic struct {
	Account     string `bson:"account" form:"account"  json:"account"`
	Password    string `bson:"password" form:"password"  json:"password"`
	Username    string `bson:"username" form:"username"  json:"username"`
	Avatar      string `bson:"avatar" form:"updatedAt" json:"avatar"`
	Role        string `bson:"role" form:"role"`
	Gender      string `bson:"gender" form:"gender"`
	CreatedTime int64  `json:"createdTime" bson:"created_time"`
	UpdatedTime int64  `json:"updatedTime" bson:"updated_time"`
}

type UserAdminBasic struct {
	Account     string `bson:"account" form:"account"  json:"account"`
	Password    string `bson:"password" form:"password"  json:"password"`
	Username    string `bson:"username" form:"username"  json:"username"`
	Avatar      string `bson:"avatar" form:"updatedAt" json:"avatar"`
	Role        string `bson:"role" form:"role"`
	Gender      string `bson:"gender" form:"gender"`
	CreatedTime int64  `json:"createdTime" bson:"created_time"`
	UpdatedTime int64  `json:"updatedTime" bson:"updated_time"`
}

type Claim struct {
	Account string `json:"account"`
	Role    string `json:"role"`
	jwt.StandardClaims
}

func (UserStudentBasic) Collection() string {
	return "user_students"
}

func (UserTeacherBasic) Collection() string {
	return "user_teachers"
}

func (UserAdminBasic) Collection() string {
	return "admin_students"
}

func (UserBasic) Collection() string {
	return "users"
}
