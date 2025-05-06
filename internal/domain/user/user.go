package user

//go:generate go run ../../../dev/gen/query_gen.go user
//go:generate go run ../../../dev/gen/builder.go user

import (
	"STUOJ/internal/infrastructure/repository/dao"
	"STUOJ/pkg/errors"
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

func (u *User) Create() (int64, error) {
	var err error

	u.CreateTime = time.Now()
	u.UpdateTime = time.Now()
	if err := u.verify(); err != nil {
		return 0, errors.ErrValidation.WithMessage(err.Error())
	}

	// 加密
	u.Password, err = u.Password.Hash()
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}

	user, err := dao.UserStore.Insert(u.toEntity())
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return int64(user.Id), nil
}

func (u *User) Update() error {
	var err error
	options := u.toOption()
	_, err = dao.UserStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}

	u.UpdateTime = time.Now()
	if err := u.verify(); err != nil {
		return errors.ErrValidation.WithMessage(err.Error())
	}

	// 加密
	u.Password, err = u.Password.Hash()
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}

	_, err = dao.UserStore.Updates(u.toEntity(), options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return nil
}

func (u *User) Delete() error {
	options := u.toOption()
	_, err := dao.UserStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	err = dao.UserStore.Delete(options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return nil
}

func WithPasswordPlaintext(password string) Option {
	return func(user *User) {
		user.Password = valueobject.NewPasswordPlaintext(password)
	}
}
