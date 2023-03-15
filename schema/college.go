package schema

type College struct {
	Description string `bson:"description" json:"description" form:"description"`
	Name        string `bson:"name" json:"name" form:"name"`
}

type UpdateCollege struct {
	Description string `bson:"description" json:"description" form:"description"`
	OldName     string `bson:"name" json:"oldName" form:"oldName"`
	NewName     string `bson:"name" json:"newName" form:"newName"`
}

func (College) Collection() string {
	return "college"
}
