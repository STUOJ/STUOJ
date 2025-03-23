package entity

//go:generate go run ../../../utils/gen/field_select.go -struct=Solution
// 题解
type Solution struct {
	Id         uint64   `gorm:"primaryKey;autoIncrement;comment:题解ID"`
	LanguageId uint64   `gorm:"not null;default:0;comment:语言ID"`
	ProblemId  uint64   `gorm:"not null;default:0;comment:题目ID"`
	SourceCode string   `gorm:"type:longtext;not null;comment:源代码"`
	Language   Language `gorm:"foreignKey:LanguageId;references:Id;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	Problem    Problem  `gorm:"foreignKey:ProblemId;references:Id;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE"`
}

func (Solution) TableName() string {
	return "tbl_solution"
}
