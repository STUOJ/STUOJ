package entity

import (
	"time"
)

// ContestStatus 比赛状态: 1 作废, 2 隐藏, 3 公开
//
//go:generate go run ../../../../dev/gen/enum_valid.go ContestStatus
type ContestStatus uint8

const (
	ContestInvalid ContestStatus = 1
	ContestHidden  ContestStatus = 2
	ContestPublic  ContestStatus = 3
)

func (s ContestStatus) String() string {
	switch s {
	case ContestInvalid:
		return "作废"
	case ContestHidden:
		return "隐藏"
	case ContestPublic:
		return "公开"
	default:
		return "未知"
	}
}

// ContestFormat 比赛赛制: 1 ACM, 2 OI, 3 IOI
//
//go:generate go run ../../../../dev/gen/enum_valid.go ContestFormat
type ContestFormat uint8

const (
	ContestACM ContestFormat = 1
	ContestOI  ContestFormat = 2
	ContestIOI ContestFormat = 3
)

func (f ContestFormat) String() string {
	switch f {
	case ContestACM:
		return "ACM"
	case ContestOI:
		return "OI"
	case ContestIOI:
		return "IOI"
	default:
		return "未知"
	}
}

// Contest 比赛
//
//go:generate go run ../../../../dev/gen/dao_store.go -struct=Contest
//go:generate go run ../../../../dev/gen/field_select.go -struct=Contest
type Contest struct {
	Id          uint64        `gorm:"primaryKey;autoIncrement;comment:比赛Id"`
	Title       string        `gorm:"type:varchar(255);not null;comment:比赛标题"`
	Description string        `gorm:"type:longtext;not null;comment:比赛描述"`
	UserId      uint64        `gorm:"not null;default:0;comment:用户Id"`
	Status      ContestStatus `gorm:"not null;default:1;comment:状态"`
	Format      ContestFormat `gorm:"not null;default:1;comment:赛制"`
	TeamSize    uint8         `gorm:"not null;default:1;comment:组队人数"`
	StartTime   time.Time     `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:开始时间"`
	EndTime     time.Time     `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:结束时间"`
	CreateTime  time.Time     `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime  time.Time     `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`
	User        User          `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Contest) TableName() string {
	return "tbl_contest"
}

// ContestUser 比赛用户关联
//
//go:generate go run ../../../../dev/gen/dao_store.go -struct=ContestUser
//go:generate go run ../../../../dev/gen/field_select.go -struct=ContestUser
type ContestUser struct {
	ContestId uint64  `gorm:"not null;default:0;primaryKey;comment:比赛Id"`
	UserId    uint64  `gorm:"not null;default:0;primaryKey;comment:用户Id"`
	User      User    `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Contest   Contest `gorm:"foreignKey:ContestId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (ContestUser) TableName() string {
	return "tbl_contest_user"
}

// ContestProblem 比赛题目关联
//
//go:generate go run ../../../../dev/gen/dao_store.go -struct=ContestProblem
//go:generate go run ../../../../dev/gen/field_select.go -struct=ContestProblem
type ContestProblem struct {
	ContestId uint64  `gorm:"not null;default:0;primaryKey;comment:比赛Id"`
	ProblemId uint64  `gorm:"not null;default:0;primaryKey;comment:题目Id"`
	Serial    uint16  `gorm:"not null;default:0;comment:排序序号"`
	Contest   Contest `gorm:"foreignKey:ContestId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Problem   Problem `gorm:"foreignKey:ProblemId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (ContestProblem) TableName() string {
	return "tbl_contest_problem"
}
