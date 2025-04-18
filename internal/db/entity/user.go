package entity

import (
	"errors"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Role 角色：0 访客，1 用户，2 编辑，3 管理，4 站长
//
//go:generate go run ../../../utils/gen/enum_valid.go Role
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

type Email string

func (e Email) Verify() error {
	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`

	// 编译正则表达式
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(string(e)) {
		return errors.New("邮箱格式不正确！")
	}
	return nil
}

func (e Email) String() string {
	return string(e)
}

// User 用户
//
//go:generate go run ../../../utils/gen/dao_store.go -struct=User
//go:generate go run ../../../utils/gen/field_select.go -struct=User
type User struct {
	Id         uint64    `gorm:"primaryKey;autoIncrement;comment:用户Id"`
	Username   string    `gorm:"type:varchar(255);not null;comment:用户名;unique"`
	Password   string    `gorm:"type:varchar(255);not null;default:123456;comment:密码"`
	Role       Role      `gorm:"not null;default:1;comment:角色"`
	Email      Email     `gorm:"type:varchar(255);not null;comment:邮箱;unique"`
	Avatar     string    `gorm:"type:text;not null;comment:头像URL"`
	Signature  string    `gorm:"type:text;not null;comment:个性签名"`
	CreateTime time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`
}

func (User) TableName() string {
	return "tbl_user"
}

// 对密码进行哈希
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// 验证密码
func (u *User) VerifyByPassword(pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pw))
}

// 验证密码
func (u *User) VerifyByHashedPassword(hpw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hpw), []byte(u.Password))
}
