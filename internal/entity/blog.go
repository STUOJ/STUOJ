package entity

import "time"

// BlogStatus 状态：1 被屏蔽, 2 草稿, 3 待审核, 4 公开, 5 公告
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

// Blog 博客
type Blog struct {
	ID         uint64     `gorm:"primaryKey;autoIncrement;comment:博客ID" json:"id"`
	UserID     uint64     `gorm:"not null;default:0;comment:用户ID" json:"user_id"`
	ProblemID  uint64     `gorm:"not null;default:0;comment:关联题目ID" json:"problem_id"`
	Title      string     `gorm:"type:text;not null;comment:标题" json:"title"`
	Content    string     `gorm:"type:longtext;not null;comment:内容" json:"content"`
	Status     BlogStatus `gorm:"not null;default:1;comment:状态" json:"status"`
	CreateTime time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`
	UpdateTime time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"`
	User       User       `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT" json:"user"`
}

func (Blog) TableName() string {
	return "tbl_blog"
}
