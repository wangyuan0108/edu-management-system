package schema

type Message struct {
	Message      string `json:"message"`
	RoomIdentity string `json:"roomIdentity"`
}

type RoomBasic struct {
	UserIdentity string `json:"userIdentity" bson:"user_identity"` // 创建者/房主/群主
	Avatar       string `json:"avatar" bson:"avatar"`              //房间头像/群头像
	Info         string `json:"info" bson:"info"`                  // 房间简介/群介绍
	Name         string `json:"name" bson:"name"`                  // 房间名/群名
	Number       string `json:"number" bson:"number"`              // 房间号/群号
	CreatedTime  int64  `json:"createdTime" bson:"created_time"`
	UpdatedTime  int64  `json:"updatedTime" bson:"updated_time"`
}

type MessageBasic struct {
	UserIdentity string `json:"userIdentity" bson:"user_identity"`
	RoomIdentity string `json:"roomIdentity" bson:"room_identity"`
	Data         string `json:"data" bson:"data"`
	CreatedTime  int64  `json:"createdTime" bson:"created_time"`
	UpdatedTime  int64  `json:"updatedTime" bson:"updated_time"`
}

func (MessageBasic) Collection() string {
	return "message_basic"
}

func (RoomBasic) Collection() string {
	return "room_basic"
}
