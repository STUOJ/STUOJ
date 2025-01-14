package entity

import "time"

// CommentStatus 状态：1 屏蔽, 2 公开
type CommentStatus uint8

const (
	CommentBanned CommentStatus = 1
	CommentPublic CommentStatus = 2
)

func (s CommentStatus) String() string {
	switch s {
	case CommentBanned:
		return "屏蔽"
	case CommentPublic:
		return "公开"
	default:
		return "未知"
	}
}

// Comment 评论
type Comment struct {
	Id         uint64        `gorm:"primaryKey;autoIncrement;comment:评论ID" json:"id"`
	UserId     uint64        `gorm:"not null;default:0;comment:用户ID" json:"-"`
	BlogId     uint64        `gorm:"not null;default:0;comment:博客ID" json:"-"`
	Content    string        `gorm:"type:longtext;not null;comment:内容" json:"content"`
	Status     CommentStatus `gorm:"not null;default:1;comment:状态" json:"status"`
	CreateTime time.Time     `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`
	UpdateTime time.Time     `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"`
	User       User          `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT" json:"user"`
	Blog       Blog          `gorm:"foreignKey:BlogId;references:Id;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE" json:"blog"`
}

func (Comment) TableName() string {
	return "tbl_comment"
}
