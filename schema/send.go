package schema

type Status struct {
	Code    int `json:"code"`
	Message any `json:"message"`
	Body    any `json:"body"`
}
