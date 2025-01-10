package entity

import "time"

// CommentStatus 状态：1 屏蔽, 2 公开
type CommentStatus uint64

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

// Comment 评论
type Comment struct {
	Id         uint64        `gorm:"primaryKey;autoIncrement;comment:评论ID" json:"id"`
	UserId     uint64        `gorm:"not null;default:0;comment:用户ID" json:"user_id"`
	BlogId     uint64        `gorm:"not null;default:0;comment:博客ID" json:"blog_id"`
	Content    string        `gorm:"type:longtext;not null;comment:内容" json:"content"`
	Status     CommentStatus `gorm:"not null;default:1;comment:状态" json:"status"`
	CreateTime time.Time     `gorm:"not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`
	UpdateTime time.Time     `gorm:"not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"`
	User       User          `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT" json:"user"`
	Blog       Blog          `gorm:"foreignKey:BlogId;references:Id;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE" json:"blog"`
}

func (Comment) TableName() string {
	return "tbl_comment"
}
