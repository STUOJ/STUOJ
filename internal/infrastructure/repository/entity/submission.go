package entity

import (
	"time"
)

// Submission 提交信息
//
//go:generate go run ../../../../dev/gen/dao_store.go -struct=Submission
//go:generate go run ../../../../dev/gen/field_select.go -struct=Submission
type Submission struct {
	Id         uint64      `gorm:"primaryKey;autoIncrement;comment:提交记录Id"`
	UserId     uint64      `gorm:"not null;default:0;comment:用户Id"`
	ProblemId  uint64      `gorm:"not null;default:0;comment:题目Id"`
	Status     JudgeStatus `gorm:"not null;default:1;comment:状态"`
	Score      uint8       `gorm:"not null;default:0;comment:分数"`
	Memory     uint64      `gorm:"not null;default:0;comment:内存（kb）"`
	Time       float64     `gorm:"not null;default:0;comment:运行耗时（s）"`
	Length     uint32      `gorm:"not null;default:0;comment:源代码长度"`
	LanguageId uint64      `gorm:"not null;default:0;comment:语言Id"`
	SourceCode string      `gorm:"type:longtext;not null;comment:源代码"`
	CreateTime time.Time   `gorm:"autoCreateTime;comment:创建时间"`
	UpdateTime time.Time   `gorm:"autoUpdateTime;comment:更新时间"`
	User       User        `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	Problem    Problem     `gorm:"foreignKey:ProblemId;references:Id;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	Language   Language    `gorm:"foreignKey:LanguageId;references:Id;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (Submission) TableName() string {
	return "tbl_submission"
}
