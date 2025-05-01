package user

//go:generate go run ../../../dev/gen/dto_gen.go user
//go:generate go run ../../../dev/gen/query_gen.go user
//go:generate go run ../../../dev/gen/builder.go user

import (
	"STUOJ/internal/infrastructure/repository/dao"
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/infrastructure/repository/entity/field"
	"STUOJ/internal/model/option"
	"STUOJ/pkg/errors"
	"time"

	"STUOJ/internal/domain/user/valueobject"
)

type User struct {
	Id         int64
	Username   valueobject.Username
	Password   valueobject.Password
	Role       entity.Role
	Email      valueobject.Email
	Avatar     valueobject.Avatar
	Signature  valueobject.Signature
	CreateTime time.Time
	UpdateTime time.Time
}

func (u *User) verify() error {
	if err := u.Username.Verify(); err != nil {
		return err
	}
	if err := u.Password.Verify(); err != nil {
		return err
	}
	if err := u.Email.Verify(); err != nil {
		return err
	}
	if err := u.Avatar.Verify(); err != nil {
		return err
	}
	if err := u.Signature.Verify(); err != nil {
		return err
	}
	return nil
}

func (u *User) toEntity() entity.User {
	return entity.User{
		Id:         uint64(u.Id),
		Username:   u.Username.String(),
		Password:   u.Password.String(),
		Role:       u.Role,
		Email:      entity.Email(u.Email.String()),
		Avatar:     u.Avatar.String(),
		Signature:  u.Signature.String(),
		CreateTime: u.CreateTime,
		UpdateTime: u.UpdateTime,
	}
}

func (u *User) fromEntity(user entity.User) *User {
	u.Id = int64(user.Id)
	u.Username = valueobject.NewUsername(user.Username)
	u.Password = valueobject.NewPassword(user.Password)
	u.Role = user.Role
	u.Email = valueobject.NewEmail(string(user.Email))
	u.Avatar = valueobject.NewAvatar(user.Avatar)
	u.Signature = valueobject.NewSignature(user.Signature)
	u.CreateTime = user.CreateTime
	u.UpdateTime = user.UpdateTime
	return u
}

func (u *User) toOption() *option.QueryOptions {
	options := option.NewQueryOptions()
	options.Filters.Add(field.UserId, option.OpEqual, u.Id)
	return options
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
