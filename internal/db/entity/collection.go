package entity

import (
	"time"
)

// CollectionStatus 题单状态: 1 私有, 2 公开
//
//go:generate go run ../../../utils/gen/enum_valid.go CollectionStatus
type CollectionStatus uint8

const (
	CollectionPrivate CollectionStatus = 1
	CollectionPublic  CollectionStatus = 2
)

func (s CollectionStatus) String() string {
	switch s {
	case CollectionPrivate:
		return "私有"
	case CollectionPublic:
		return "公开"
	default:
		return "未知"
	}
}

//go:generate go run ../../../utils/gen/dao_store.go -struct=Collection
//go:generate go run ../../../utils/gen/field_select.go -struct=Collection
type Collection struct {
	Id          uint64           `gorm:"primaryKey;autoIncrement;comment:题单ID"`
	UserId      uint64           `gorm:"not null;default:0;comment:用户ID"`
	Title       string           `gorm:"type:text;not null;comment:标题"`
	Description string           `gorm:"type:longtext;not null;comment:简介"`
	Status      CollectionStatus `gorm:"not null;default:1;comment:状态"`
	CreateTime  time.Time        `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime  time.Time        `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`
	User        User             `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Collection) TableName() string {
	return "tbl_collection"
}

//go:generate go run ../../../utils/gen/dao_store.go -struct=CollectionUser
//go:generate go run ../../../utils/gen/field_select.go -struct=CollectionUser
type CollectionUser struct {
	CollectionId uint64     `gorm:"not null;default:0;primaryKey;comment:题单ID"`
	UserId       uint64     `gorm:"not null;default:0;primaryKey;comment:用户ID"`
	User         User       `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Collection   Collection `gorm:"foreignKey:CollectionId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (CollectionUser) TableName() string {
	return "tbl_collection_user"
}

//go:generate go run ../../../utils/gen/dao_store.go -struct=CollectionProblem
//go:generate go run ../../../utils/gen/field_select.go -struct=CollectionProblem
type CollectionProblem struct {
	CollectionId uint64     `gorm:"not null;default:0;primaryKey;comment:题单ID"`
	ProblemId    uint64     `gorm:"not null;default:0;primaryKey;comment:题目ID"`
	Serial       uint16     `gorm:"not null;default:0;comment:排序序号"`
	Collection   Collection `gorm:"foreignKey:CollectionId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Problem      Problem    `gorm:"foreignKey:ProblemId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (CollectionProblem) TableName() string {
	return "tbl_collection_problem"
}
