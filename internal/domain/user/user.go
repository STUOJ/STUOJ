package user

//go:generate go run ../../../dev/gen/query_gen.go user
//go:generate go run ../../../dev/gen/builder.go user

import (
	"time"

	"STUOJ/internal/domain/user/valueobject"
)

type User struct {
	Id         valueobject.Id
	Username   valueobject.Username
	Password   valueobject.Password
	Role       valueobject.Role
	Email      valueobject.Email
	Avatar     valueobject.Avatar
	Signature  valueobject.Signature
	CreateTime time.Time
	UpdateTime time.Time
}

func WithPasswordPlaintext(password string) Option {
	return func(user *User) {
		user.Password = valueobject.NewPasswordPlaintext(password)
	}
}
