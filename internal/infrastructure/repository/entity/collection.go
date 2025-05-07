package entity

import (
	"time"
)

// CollectionStatus 题单状态: 1 删除, 2 私有, 3 公开
//
//go:generate go run ../../../../dev/gen/enum_valid.go CollectionStatus
type CollectionStatus uint8

const (
	CollectionDeleted CollectionStatus = 1
	CollectionPrivate CollectionStatus = 2
	CollectionPublic  CollectionStatus = 3
)

func (s CollectionStatus) String() string {
	switch s {
	case CollectionDeleted:
		return "删除"
	case CollectionPrivate:
		return "私有"
	case CollectionPublic:
		return "公开"
	default:
		return "未知"
	}
}

//go:generate go run ../../../../dev/gen/dao_store.go -struct=Collection
//go:generate go run ../../../../dev/gen/field_select.go -struct=Collection
type Collection struct {
	Id          uint64           `gorm:"primaryKey;autoIncrement;comment:题单Id"`
	UserId      uint64           `gorm:"not null;default:0;comment:用户Id"`
	Title       string           `gorm:"type:text;not null;comment:标题"`
	Description string           `gorm:"type:longtext;not null;comment:简介"`
	Status      CollectionStatus `gorm:"not null;default:1;comment:状态"`
	CreateTime  time.Time        `gorm:"autoCreateTime;comment:创建时间"`
	UpdateTime  time.Time        `gorm:"autoUpdateTime;comment:更新时间"`
	User        User             `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Collection) TableName() string {
	return "tbl_collection"
}

//go:generate go run ../../../../dev/gen/dao_store.go -struct=CollectionUser
//go:generate go run ../../../../dev/gen/field_select.go -struct=CollectionUser
type CollectionUser struct {
	CollectionId uint64     `gorm:"not null;default:0;primaryKey;comment:题单Id"`
	UserId       uint64     `gorm:"not null;default:0;primaryKey;comment:用户Id"`
	User         User       `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Collection   Collection `gorm:"foreignKey:CollectionId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (CollectionUser) TableName() string {
	return "tbl_collection_user"
}

//go:generate go run ../../../../dev/gen/dao_store.go -struct=CollectionProblem
//go:generate go run ../../../../dev/gen/field_select.go -struct=CollectionProblem
type CollectionProblem struct {
	CollectionId uint64     `gorm:"not null;default:0;primaryKey;comment:题单Id"`
	ProblemId    uint64     `gorm:"not null;default:0;primaryKey;comment:题目Id"`
	Serial       uint16     `gorm:"not null;default:0;comment:排序序号"`
	Collection   Collection `gorm:"foreignKey:CollectionId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Problem      Problem    `gorm:"foreignKey:ProblemId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (CollectionProblem) TableName() string {
	return "tbl_collection_problem"
}
