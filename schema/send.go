package schema

type Status struct {
	Code    int    `json:"code"`    // 业务码
	Message string `json:"message"` // 响应消息
	Body    any    `json:"body"`    // 消息体
}
