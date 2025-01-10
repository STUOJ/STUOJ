package entity

// 题解
type Solution struct {
	Id         uint64   `gorm:"primaryKey;autoIncrement;comment:题解ID" json:"id"`
	LanguageId uint64   `gorm:"not null;default:0;comment:语言ID" json:"language_id"`
	ProblemId  uint64   `gorm:"not null;default:0;comment:题目ID" json:"problem_id"`
	SourceCode string   `gorm:"type:longtext;not null;comment:源代码" json:"source_code"`
	Language   Language `gorm:"foreignKey:LanguageId;references:Id;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT" json:"-"`
	Problem    Problem  `gorm:"foreignKey:ProblemId;references:Id;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE" json:"-"`
}

func (Solution) TableName() string {
	return "tbl_solution"
}
