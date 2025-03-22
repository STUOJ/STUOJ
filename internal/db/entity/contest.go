package entity

import (
	"time"
)

// ContestStatus 比赛状态: 1 作废, 2 编辑, 3 准备, 4 进行, 5 结束
type ContestStatus uint8

const (
	ContestInvalid    ContestStatus = 1
	ContestEditing    ContestStatus = 2
	ContestReady      ContestStatus = 3
	ContestProcessing ContestStatus = 4
	ContestEnded      ContestStatus = 5
)

func (s ContestStatus) String() string {
	switch s {
	case ContestInvalid:
		return "作废"
	case ContestEditing:
		return "编辑"
	case ContestReady:
		return "准备"
	case ContestProcessing:
		return "进行"
	case ContestEnded:
		return "结束"
	default:
		return "未知"
	}
}

// ContestFormat 比赛赛制: 1 ACM, 2 OI, 3 IOI
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
type Contest struct {
	Id           uint64        `gorm:"primaryKey;autoIncrement;comment:比赛ID"`
	UserId       uint64        `gorm:"not null;default:0;comment:用户ID"`
	CollectionId uint64        `gorm:"not null;default:0;comment:题单ID"`
	Status       ContestStatus `gorm:"not null;default:1;comment:状态"`
	Format       ContestFormat `gorm:"not null;default:1;comment:赛制"`
	TeamSize     uint8         `gorm:"not null;default:1;comment:组队人数"`
	StartTime    time.Time     `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:开始时间"`
	EndTime      time.Time     `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:结束时间"`
	CreateTime   time.Time     `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime   time.Time     `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`
	User         User          `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Collection   Collection    `gorm:"foreignKey:CollectionId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Contest) TableName() string {
	return "tbl_contest"
}
