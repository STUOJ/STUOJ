package entity

// 题解
//
//go:generate go run ../../../utils/gen/dao_store.go -struct=Solution
//go:generate go run ../../../utils/gen/field_select.go -struct=Solution
type Solution struct {
	Id         uint64   `gorm:"primaryKey;autoIncrement;comment:题解Id"`
	LanguageId uint64   `gorm:"not null;default:0;comment:语言Id"`
	ProblemId  uint64   `gorm:"not null;default:0;comment:题目Id"`
	SourceCode string   `gorm:"type:longtext;not null;comment:源代码"`
	Language   Language `gorm:"foreignKey:LanguageId;references:Id;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	Problem    Problem  `gorm:"foreignKey:ProblemId;references:Id;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE"`
}

func (Solution) TableName() string {
	return "tbl_solution"
}
