package entity

import (
	"time"
)

// Submission 提交信息
type Submission struct {
	ID         uint64      `gorm:"primaryKey;autoIncrement;comment:提交记录ID" json:"id"`
	UserID     uint64      `gorm:"not null;default:0;comment:用户ID" json:"user_id"`
	ProblemID  uint64      `gorm:"not null;default:0;comment:题目ID" json:"problem_id"`
	Status     JudgeStatus `gorm:"not null;default:1;comment:状态" json:"status"`
	Score      uint64      `gorm:"not null;default:0;comment:分数" json:"score"`
	LanguageID uint64      `gorm:"not null;default:0;comment:语言ID" json:"language_id"`
	Length     uint64      `gorm:"not null;default:0;comment:源代码长度" json:"length"`
	Memory     uint64      `gorm:"not null;default:0;comment:内存（kb）" json:"memory"`
	Time       float64     `gorm:"not null;default:0;comment:运行耗时（s）" json:"time"`
	SourceCode string      `gorm:"type:longtext;not null;comment:源代码" json:"source_code"`
	CreateTime time.Time   `gorm:"not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`
	UpdateTime time.Time   `gorm:"not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"`
	User       User        `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT" json:"user"`
	Problem    Problem     `gorm:"foreignKey:ProblemID;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT" json:"problem"`
	Language   Language    `gorm:"foreignKey:LanguageID;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT" json:"language"`
}

func (Submission) TableName() string {
	return "tbl_submission"
}
