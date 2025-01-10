package entity

import "time"

// 1 被屏蔽, 2 草稿, 3 待审核, 4 公开, 5 公告
type BlogStatus uint8

const (
	BLogStatusBanned BlogStatus = 1
	BlogStatusDraft  BlogStatus = 2
	BLogStatusReview BlogStatus = 3
	BlogStatusPublic BlogStatus = 4
	BlogStatusNotice BlogStatus = 5
)

func (s BlogStatus) String() string {
	switch s {
	case BLogStatusBanned:
		return "屏蔽"
	case BlogStatusPublic:
		return "公开"
	case BlogStatusDraft:
		return "草稿"
	case BLogStatusReview:
		return "审核"
	case BlogStatusNotice:
		return "公告"
	default:
		return "未知"
	}
}

// 博客
type Blog struct {
	Id         uint64     `gorm:"primaryKey;autoIncrement;comment:'博客ID'" json:"id,omitempty"`
	UserId     uint64     `gorm:"not null;default:0;comment:'用户ID';constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;" json:"user_id,omitempty"`
	ProblemId  uint64     `gorm:"not null;default:0;comment:'关联题目ID'" json:"problem_id,omitempty"`
	Title      string     `gorm:"type:text;not null;comment:'标题'" json:"title,omitempty"`
	Content    string     `gorm:"type:longtext;not null;comment:'内容'" json:"content,omitempty"`
	Status     BlogStatus `gorm:"not null;default:1;comment:'状态'" json:"status"`
	CreateTime time.Time  `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:'创建时间'" json:"create_time,omitempty"`
	UpdateTime time.Time  `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP;comment:'更新时间'" json:"update_time,omitempty"`
}

func (Blog) TableName() string {
	return "tbl_blog"
}
