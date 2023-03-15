package schema

type Specialty struct {
	Id          string `bson:"_id" json:"id"`
	Code        string `bson:"code" json:"code"`
	College     string `bson:"college" json:"college"`
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
}

type SpecialtyWithCollege struct {
	OldCollegeName          string `form:"oldCollegeName" bson:"college" json:"oldCollegeName"`
	OldSpecialtyName        string `form:"oldSpecialtyName" bson:"name" json:"oldSpecialtyName"`
	OldCollegeDescription   string `form:"oldCollegeDescription" bson:"description" json:"oldCollegeDescription"`
	NewCollegeName          string `form:"newCollegeName" bson:"college" json:"newCollegeName"`
	NewSpecialtyName        string `form:"newCollegeDescription" bson:"description" json:"newCollegeDescription"`
	NewSpecialtyDescription string `form:"newSpecialtyName" bson:"name" json:"newSpecialtyName"`
}

func (Specialty) Collection() string {
	return "specialty"
}
