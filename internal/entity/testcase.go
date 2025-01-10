package entity

// Testcase 测试用例
type Testcase struct {
	ID         uint64  `gorm:"primaryKey;autoIncrement;comment:评测点ID" json:"id"`
	Serial     uint64  `gorm:"not null;default:0;comment:评测点序号" json:"serial"`
	ProblemID  uint64  `gorm:"not null;default:0;comment:题目ID" json:"problem_id"`
	TestInput  string  `gorm:"type:longtext;not null;comment:测试输入" json:"test_input"`
	TestOutput string  `gorm:"type:longtext;not null;comment:测试输出" json:"test_output"`
	Problem    Problem `gorm:"foreignKey:ProblemID;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE" json:"problem"`
}

func (Testcase) TableName() string {
	return "tbl_testcase"
}
