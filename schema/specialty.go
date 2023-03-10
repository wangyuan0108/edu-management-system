package schema

type Specialty struct {
	Code    string `bson:"code" json:"code"`
	College string `bson:"college" json:"college"`
	Name    string `bson:"name" json:"name"`
}

func (Specialty) Collection() string {
	return "specialty"
}
