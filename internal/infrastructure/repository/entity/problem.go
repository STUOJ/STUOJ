package entity

import (
	"time"
)

// Difficulty 难度：0 未知，1 入门，2 简单，3 中等，4 较难，5 困难，6 超难
//
//go:generate go run ../../../../dev/gen/enum_valid.go Difficulty
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

// ProblemStatus 题目状态: 1 删除, 2 出题, 3 调试, 4 隐藏, 5 公开
type ProblemStatus uint64

const (
	ProblemDeleted   ProblemStatus = 1
	ProblemEditing   ProblemStatus = 2
	ProblemDebugging ProblemStatus = 3
	ProblemHidden    ProblemStatus = 4
	ProblemPublic    ProblemStatus = 5
)

func (s ProblemStatus) String() string {
	switch s {
	case ProblemDeleted:
		return "删除"
	case ProblemEditing:
		return "出题"
	case ProblemDebugging:
		return "调试"
	case ProblemHidden:
		return "隐藏"
	case ProblemPublic:
		return "公开"
	default:
		return "未知"
	}
}

// Problem 题目
//
//go:generate go run ../../../../dev/gen/dao_store.go -struct=Problem
//go:generate go run ../../../../dev/gen/field_select.go -struct=Problem
type Problem struct {
	Id           uint64        `gorm:"primaryKey;autoIncrement;comment:题目Id"`
	Title        string        `gorm:"type:text;not null;comment:标题"`
	Source       string        `gorm:"type:text;not null;comment:题目来源"`
	Difficulty   Difficulty    `gorm:"not null;default:0;comment:难度"`
	TimeLimit    float64       `gorm:"not null;default:1;comment:时间限制（s）"`
	MemoryLimit  uint64        `gorm:"not null;default:131072;comment:内存限制（kb）"`
	Description  string        `gorm:"type:longtext;not null;comment:题面"`
	Input        string        `gorm:"type:longtext;not null;comment:输入说明"`
	Output       string        `gorm:"type:longtext;not null;comment:输出说明"`
	SampleInput  string        `gorm:"type:longtext;not null;comment:输入样例"`
	SampleOutput string        `gorm:"type:longtext;not null;comment:输出样例"`
	Hint         string        `gorm:"type:longtext;not null;comment:提示"`
	Status       ProblemStatus `gorm:"not null;default:1;comment:状态"`
	CreateTime   time.Time     `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime   time.Time     `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`
}

func (Problem) TableName() string {
	return "tbl_problem"
}

// 题目标签关系
//
//go:generate go run ../../../../dev/gen/dao_store.go -struct=ProblemTag
//go:generate go run ../../../../dev/gen/field_select.go -struct=ProblemTag
type ProblemTag struct {
	ProblemId uint64  `gorm:"primaryKey;not null;default:0;comment:题目Id"`
	TagId     uint64  `gorm:"primaryKey;not null;default:0;comment:标签Id"`
	Problem   Problem `gorm:"foreignKey:ProblemId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tag       Tag     `gorm:"foreignKey:TagId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (ProblemTag) TableName() string {
	return "tbl_problem_tag"
}
