package entity

import "time"

// 状态：1 被屏蔽, 2 公开
type CommentStatus uint8

const (
	CommentStatusBanned CommentStatus = 1
	CommentStatusPublic CommentStatus = 2
)

func (s CommentStatus) String() string {
	switch s {
	case CommentStatusBanned:
		return "屏蔽"
	case CommentStatusPublic:
		return "公开"
	default:
		return "未知"
	}
}

// 评论
type Comment struct {
	Id         uint64        `gorm:"primaryKey;autoIncrement;comment:'评论ID'" json:"id,omitempty"`
	UserId     uint64        `gorm:"not null;default:0;comment:'用户ID'" json:"user_id,omitempty"`
	BlogId     uint64        `gorm:"not null;default:0;comment:'博客ID'" json:"blog_id,omitempty"`
	Content    string        `gorm:"type:longtext;not null;comment:'内容'" json:"content,omitempty"`
	Status     CommentStatus `gorm:"not null;default:1;comment:'状态：0 被屏蔽，1 公开'" json:"status"`
	CreateTime time.Time     `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:'创建时间'" json:"create_time,omitempty"`
	UpdateTime time.Time     `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP;comment:'更新时间'" json:"update_time,omitempty"`
}

func (Comment) TableName() string {
	return "tbl_comment"
}
