package user

import (
	"STUOJ/internal/db/entity"
	"STUOJ/internal/domain/user/valueobject"
	"time"
)

type Option func(*User)

func NewUser(options ...Option) *User {
	user := &User{}
	for _, option := range options {
		option(user)
	}
	return user
}

func WithId(id uint64) Option {
	return func(u *User) {
		u.Id = id
	}
}

func WithUsername(username string) Option {
	return func(u *User) {
		u.Username = valueobject.NewUsername(username)
	}
}

func WithPassword(password string) Option {
	return func(u *User) {
		u.Password = valueobject.NewPassword(password)
	}
}

func WithRole(role entity.Role) Option {
	return func(u *User) {
		u.Role = role
	}
}

func WithEmail(email string) Option {
	return func(u *User) {
		u.Email = valueobject.NewEmail(email)
	}
}

func WithAvatar(avatar string) Option {
	return func(u *User) {
		u.Avatar = valueobject.NewAvatar(avatar)
	}
}

func WithSignature(signature string) Option {
	return func(u *User) {
		u.Signature = valueobject.NewSignature(signature)
	}
}

func WithCreateTime(createTime time.Time) Option {
	return func(u *User) {
		u.CreateTime = createTime
	}
}

func WithUpdateTime(updateTime time.Time) Option {
	return func(u *User) {
		u.UpdateTime = updateTime
	}
}
