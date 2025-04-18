package user

//go:generate go run ../../../utils/gen/dto_gen.go user
//go:generate go run ../../../utils/gen/query_gen.go user

import (
	"time"

	"STUOJ/internal/db/dao"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/domain/user/valueobject"
	"STUOJ/internal/errors"
)

type User struct {
	Id         uint64
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
		Id:         u.Id,
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
	u.Id = user.Id
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

func (u *User) Create() (uint64, error) {
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
	return user.Id, &errors.NoError
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
	return &errors.NoError
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
	return &errors.NoError
}
