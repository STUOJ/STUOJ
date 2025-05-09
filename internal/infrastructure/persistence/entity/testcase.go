package entity

import "time"

// Testcase 测试用例
//
//go:generate go run ../../../../dev/gen/dao_store.go -struct=Testcase
//go:generate go run ../../../../dev/gen/field_select.go -struct=Testcase
type Testcase struct {
	Id         uint64    `gorm:"primaryKey;autoIncrement;comment:评测点Id"`
	ProblemId  uint64    `gorm:"not null;default:0;comment:题目Id"`
	Serial     uint16    `gorm:"not null;default:0;comment:评测点序号"`
	TestInput  string    `gorm:"type:longtext;not null;comment:测试输入"`
	TestOutput string    `gorm:"type:longtext;not null;comment:测试输出"`
	CreateTime time.Time `gorm:"autoCreateTime;comment:创建时间"`
	UpdateTime time.Time `gorm:"autoUpdateTime;comment:更新时间"`
	Problem    Problem   `gorm:"foreignKey:ProblemId;references:id;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE"`
}

func (Testcase) TableName() string {
	return "tbl_testcase"
}
