package entity

import "time"

// CommentStatus 状态：1 删除, 2 公开
//
//go:generate go run ../../../../dev/gen/enum_valid.go
type CommentStatus uint8

const (
	CommentDeleted CommentStatus = 1
	CommentPublic  CommentStatus = 2
)

func (s CommentStatus) String() string {
	switch s {
	case CommentDeleted:
		return "删除"
	case CommentPublic:
		return "公开"
	default:
		return "未知"
	}
}

// Comment 评论
//
//go:generate go run ../../../../dev/gen/dao_store.go -struct=Comment
//go:generate go run ../../../../dev/gen/field_select.go -struct=Comment
type Comment struct {
	Id         uint64        `gorm:"primaryKey;autoIncrement;comment:评论Id"`
	UserId     uint64        `gorm:"not null;default:0;comment:用户Id"`
	BlogId     uint64        `gorm:"not null;default:0;comment:博客Id"`
	Content    string        `gorm:"type:longtext;not null;comment:内容"`
	Status     CommentStatus `gorm:"not null;default:1;comment:状态"`
	CreateTime time.Time     `gorm:"autoCreateTime;comment:创建时间"`
	UpdateTime time.Time     `gorm:"autoUpdateTime;comment:更新时间"`
	User       User          `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	Blog       Blog          `gorm:"foreignKey:BlogId;references:Id;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE"`
}

func (Comment) TableName() string {
	return "tbl_comment"
}
