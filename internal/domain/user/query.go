package user

import (
	"STUOJ/internal/db/dao"
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/errors"
	"STUOJ/internal/model/querymodel"
)

type _Query struct{}

var Query = new(_Query)

func (*_Query) SelectById(id uint64) (User, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.UserId, option.OpEqual, id)
	user, err := dao.UserStore.SelectOne(options)
	if err != nil {
		return User{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewUser().fromEntity(user), &errors.NoError
}

func (*_Query) SelectSimpleById(id uint64) (User, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.UserId, option.OpEqual, id)
	options.Field = query.UserSimpleField
	user, err := dao.UserStore.SelectOne(options)
	if err != nil {
		return User{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewUser().fromEntity(user), &errors.NoError
}

func (*_Query) SelectByUsername(username string) (User, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.UserUsername, option.OpEqual, username)
	options.Field = query.UserAllField
	user, err := dao.UserStore.SelectOne(options)
	if err != nil {
		return User{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewUser().fromEntity(user), &errors.NoError
}

func (*_Query) SelectByEmail(email string) (User, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.UserEmail, option.OpEqual, email)
	options.Field = query.UserAllField
	user, err := dao.UserStore.SelectOne(options)
	if err != nil {
		return User{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewUser().fromEntity(user), &errors.NoError
}

func (*_Query) Select(model querymodel.UserQueryModel) ([]User, error) {
	queryOptions := model.GenerateOptions()
	queryOptions.Field = query.UserAllField
	users, err := dao.UserStore.Select(queryOptions)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var result []User
	for _, user := range users {
		result = append(result, *NewUser().fromEntity(user))
	}
	return result, &errors.NoError
}

func (*_Query) Count(model querymodel.UserQueryModel) (int64, error) {
	queryOptions := model.GenerateOptions()
	count, err := dao.UserStore.Count(queryOptions)
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return count, &errors.NoError
}
