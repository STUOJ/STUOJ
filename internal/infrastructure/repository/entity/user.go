package entity

import (
	"time"
)

// Role 角色：0 访客，1 用户，2 编辑，3 管理，4 站长
//
//go:generate go run ../../../../dev/gen/enum_valid.go Role
type Role uint8

const (
	RoleVisitor Role = 0
	RoleUser    Role = 1
	RoleEditor  Role = 2
	RoleAdmin   Role = 3
	RoleRoot    Role = 4
)

var roleNames = map[Role]string{
	RoleVisitor: "访客",
	RoleUser:    "用户",
	RoleEditor:  "编辑",
	RoleAdmin:   "管理",
	RoleRoot:    "站长",
}

func (r Role) String() string {
	if name, ok := roleNames[r]; ok {
		return name
	}
	return "未知"
}

// User 用户
//
//go:generate go run ../../../../dev/gen/dao_store.go -struct=User
//go:generate go run ../../../../dev/gen/field_select.go -struct=User
type User struct {
	Id         uint64    `gorm:"primaryKey;autoIncrement;comment:用户Id"`
	Username   string    `gorm:"type:varchar(255);not null;comment:用户名;unique"`
	Password   string    `gorm:"type:varchar(255);not null;default:123456;comment:密码"`
	Role       Role      `gorm:"not null;default:1;comment:角色"`
	Email      string    `gorm:"type:varchar(255);not null;comment:邮箱;unique"`
	Avatar     string    `gorm:"type:text;not null;comment:头像URL"`
	Signature  string    `gorm:"type:text;not null;comment:个性签名"`
	CreateTime time.Time `gorm:"autoCreateTime;comment:创建时间"`
	UpdateTime time.Time `gorm:"autoUpdateTime;comment:更新时间"`
}

func (User) TableName() string {
	return "tbl_user"
}
