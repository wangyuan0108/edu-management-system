package schema

type Class struct {
	Id         string `bson:"_id" json:"id"`
	Specialty  string `bson:"specialty" json:"specialty"`
	ClassName  string `bson:"class_name" json:"className"`
	Population int32  `bson:"population" json:"population"`
}

func (Class) Collection() string {
	return "class_list"
}

func (Class) Name() string {
	return "班级列表"
}
