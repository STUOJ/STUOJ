package entity

import (
	"time"
)

// Difficulty 难度：0 未知，1 入门，2 简单，3 中等，4 较难，5 困难，6 超难
type Difficulty uint8

const (
	DifficultyU Difficulty = 0
	DifficultyE Difficulty = 1
	DifficultyD Difficulty = 2
	DifficultyC Difficulty = 3
	DifficultyB Difficulty = 4
	DifficultyA Difficulty = 5
	DifficultyS Difficulty = 6
)

func (d Difficulty) String() string {
	switch d {
	case DifficultyU:
		return "[?]未知"
	case DifficultyE:
		return "[E]入门"
	case DifficultyD:
		return "[D]简单"
	case DifficultyC:
		return "[C]中等"
	case DifficultyB:
		return "[B]较难"
	case DifficultyA:
		return "[A]困难"
	case DifficultyS:
		return "[S]超难"
	default:
		return "[?]未知"
	}
}

// ProblemStatus 题目状态: 1 作废, 2 出题, 3 调试, 4 隐藏, 5 比赛, 6 会员, 7 公开
type ProblemStatus uint64

const (
	ProblemStatusInvalid   ProblemStatus = 1
	ProblemStatusEditing   ProblemStatus = 2
	ProblemStatusDebugging ProblemStatus = 3
	ProblemStatusHidden    ProblemStatus = 4
	ProblemStatusContest   ProblemStatus = 5
	ProblemStatusVip       ProblemStatus = 6
	ProblemStatusPublic    ProblemStatus = 7
)

func (s ProblemStatus) String() string {
	switch s {
	case ProblemStatusInvalid:
		return "作废"
	case ProblemStatusEditing:
		return "出题"
	case ProblemStatusDebugging:
		return "调试"
	case ProblemStatusHidden:
		return "隐藏"
	case ProblemStatusContest:
		return "比赛"
	case ProblemStatusVip:
		return "会员"
	case ProblemStatusPublic:
		return "公开"
	default:
		return "未知"
	}
}

// Problem 题目
type Problem struct {
	Id           uint64        `gorm:"primaryKey;autoIncrement;comment:题目ID" json:"id"`
	Title        string        `gorm:"type:text;not null;comment:标题" json:"title"`
	Source       string        `gorm:"type:text;not null;comment:题目来源" json:"source"`
	Difficulty   Difficulty    `gorm:"not null;default:0;comment:难度" json:"difficulty"`
	TimeLimit    float64       `gorm:"not null;default:1;comment:时间限制（s）" json:"time_limit"`
	MemoryLimit  uint64        `gorm:"not null;default:131072;comment:内存限制（kb）" json:"memory_limit"`
	Description  string        `gorm:"type:longtext;not null;comment:题面" json:"description"`
	Input        string        `gorm:"type:longtext;not null;comment:输入说明" json:"input"`
	Output       string        `gorm:"type:longtext;not null;comment:输出说明" json:"output"`
	SampleInput  string        `gorm:"type:longtext;not null;comment:输入样例" json:"sample_input"`
	SampleOutput string        `gorm:"type:longtext;not null;comment:输出样例" json:"sample_output"`
	Hint         string        `gorm:"type:longtext;not null;comment:提示" json:"hint"`
	Status       ProblemStatus `gorm:"not null;default:1;comment:状态" json:"status"`
	CreateTime   time.Time     `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`
	UpdateTime   time.Time     `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"`
	Tags         []Tag         `gorm:"many2many:tbl_problem_tag;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE" json:"tags"`
}

func (Problem) TableName() string {
	return "tbl_problem"
}
