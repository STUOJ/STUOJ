package entity

import (
	"time"
)

// CollectionStatus 题单状态: 1 私有, 2 公开
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

type Collection struct {
	Id          uint64           `gorm:"primaryKey;autoIncrement;comment:题单ID" json:"id"`
	UserId      uint64           `gorm:"not null;default:0;comment:用户ID" json:"user_id"`
	Title       string           `gorm:"type:text;not null;comment:标题" json:"title"`
	Description string           `gorm:"type:longtext;not null;comment:简介" json:"description"`
	Status      CollectionStatus `gorm:"not null;default:1;comment:状态" json:"status"`
	CreateTime  time.Time        `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`
	UpdateTime  time.Time        `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"`
	User        User             `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	UserIds     []uint64         `gorm:"-" json:"user_ids"`
	ProblemIds  []uint64         `gorm:"-" json:"problem_ids"`
}

func (Collection) TableName() string {
	return "tbl_collection"
}

type CollectionUser struct {
	CollectionId uint64     `gorm:"not null;default:0;primaryKey;comment:题单ID" json:"collection_id"`
	UserId       uint64     `gorm:"not null;default:0;primaryKey;comment:用户ID" json:"user_id"`
	User         User       `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Collection   Collection `gorm:"foreignKey:CollectionId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (CollectionUser) TableName() string {
	return "tbl_collection_user"
}

type CollectionProblem struct {
	CollectionId uint64     `gorm:"not null;default:0;primaryKey;comment:题单ID" json:"collection_id"`
	ProblemId    uint64     `gorm:"not null;default:0;primaryKey;comment:题目ID" json:"problem_id"`
	Collection   Collection `gorm:"foreignKey:CollectionId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Problem      Problem    `gorm:"foreignKey:ProblemId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (CollectionProblem) TableName() string {
	return "tbl_collection_problem"
}
