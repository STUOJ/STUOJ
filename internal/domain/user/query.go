package user

import (
	"STUOJ/internal/db/dao"
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/errors"
)

type _Query struct{}

var Query = new(_Query)

func (query *_Query) SelectById(id uint64) (*User, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.UserId, option.OpEqual, id)
	user, err := dao.UserStore.SelectOne(options)
	if err != nil {
		return nil, errors.ErrNotFound.WithMessage(err.Error())
	}
	return NewUser().fromEntity(user), nil
}

func (query *_Query) SelectByUsername(username string) (*User, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.UserUsername, option.OpEqual, username)
	user, err := dao.UserStore.SelectOne(options)
	if err != nil {
		return nil, errors.ErrNotFound.WithMessage(err.Error())
	}
	return NewUser().fromEntity(user), nil
}

func (query *_Query) SelectByEmail(email string) (*User, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.UserEmail, option.OpEqual, email)
	user, err := dao.UserStore.SelectOne(options)
	if err != nil {
		return nil, errors.ErrNotFound.WithMessage(err.Error())
	}
	return NewUser().fromEntity(user), nil
}

func (query *_Query) Select(options *option.QueryOptions) ([]*User, error) {
	users, err := dao.UserStore.Select(options)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var result []*User
	for _, user := range users {
		result = append(result, NewUser().fromEntity(user))
	}
	return result, nil
}

func (query *_Query) Count(options *option.QueryOptions) (int64, error) {
	count, err := dao.UserStore.Count(options)
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return count, nil
}
