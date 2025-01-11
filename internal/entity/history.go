package entity

import (
	"time"
)

// Operation 操作：0 未知，1 添加，2 修改，3 删除
type Operation uint8

const (
	OperationUnknown Operation = 0
	OperationInsert  Operation = 1
	OperationUpdate  Operation = 2
	OperationDelete  Operation = 3
)

func (o Operation) String() string {
	switch o {
	case OperationUnknown:
		return "未知"
	case OperationInsert:
		return "添加"
	case OperationUpdate:
		return "修改"
	case OperationDelete:
		return "删除"
	default:
		return "未知"
	}
}

// History 题目历史记录
type History struct {
	Id           uint64     `gorm:"primaryKey;autoIncrement;comment:记录ID" json:"id"`
	UserId       uint64     `gorm:"not null;default:0;comment:用户ID" json:"-"`
	ProblemId    uint64     `gorm:"not null;default:0;comment:题目ID" json:"problem_id"`
	Title        string     `gorm:"type:text;not null;comment:标题" json:"title"`
	Source       string     `gorm:"type:text;not null;comment:题目来源" json:"source"`
	Difficulty   Difficulty `gorm:"not null;default:0;comment:难度" json:"difficulty"`
	TimeLimit    float64    `gorm:"not null;default:1;comment:时间限制（s）" json:"time_limit"`
	MemoryLimit  uint64     `gorm:"not null;default:131072;comment:内存限制（kb）" json:"memory_limit"`
	Description  string     `gorm:"type:longtext;not null;comment:题面" json:"description"`
	Input        string     `gorm:"type:longtext;not null;comment:输入说明" json:"input"`
	Output       string     `gorm:"type:longtext;not null;comment:输出说明" json:"output"`
	SampleInput  string     `gorm:"type:longtext;not null;comment:输入样例" json:"sample_input"`
	SampleOutput string     `gorm:"type:longtext;not null;comment:输出样例" json:"sample_output"`
	Hint         string     `gorm:"type:longtext;not null;comment:提示" json:"hint"`
	Operation    Operation  `gorm:"not null;default:0;comment:操作" json:"operation"`
	CreateTime   time.Time  `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`
	User         User       `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT" json:"user"`
}

func (History) TableName() string {
	return "tbl_history"
}
