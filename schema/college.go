package schema

type College struct {
	Info string `bson:"info" json:"info" form:"info"`
	Name string `bson:"name" json:"name" form:"name"`
}

func (College) Collection() string {
	return "college"
}
