package entity

import "time"

// BlogStatus 状态：1 屏蔽, 2 草稿, 3 公开, 4 公告
type BlogStatus uint8

const (
	BlogBanned BlogStatus = 1
	BlogDraft  BlogStatus = 2
	BlogPublic BlogStatus = 3
	BlogNotice BlogStatus = 4
)

func (s BlogStatus) String() string {
	switch s {
	case BlogBanned:
		return "屏蔽"
	case BlogPublic:
		return "公开"
	case BlogDraft:
		return "草稿"
	case BlogNotice:
		return "公告"
	default:
		return "未知"
	}
}

// Blog 博客
type Blog struct {
	Id         uint64     `gorm:"primaryKey;autoIncrement;comment:博客ID" json:"id"`
	UserId     uint64     `gorm:"not null;default:0;comment:用户ID" json:"-"`
	ProblemId  uint64     `gorm:"not null;default:0;comment:关联题目ID" json:"-"`
	Title      string     `gorm:"type:text;not null;comment:标题" json:"title"`
	Content    string     `gorm:"type:longtext;not null;comment:内容" json:"content"`
	Status     BlogStatus `gorm:"not null;default:1;comment:状态" json:"status"`
	CreateTime time.Time  `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`
	UpdateTime time.Time  `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"`
	User       User       `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT" json:"user"`
	Problem    Problem    `json:"problem,omitempty"`
}

func (Blog) TableName() string {
	return "tbl_blog"
}
