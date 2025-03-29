package contest

import (
	"STUOJ/internal/db/dao"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/errors"
)

type _Query struct{}

var Query = new(_Query)

func (_Query) Select(options *option.QueryOptions) ([]Contest, error) {
	contests, err := dao.ContestStore.Select(options)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var result []Contest
	for _, contest := range contests {
		result = append(result, *NewContest().fromEntity(contest))
	}
	return result, &errors.NoError
}

func (_Query) SelectById(id uint64) (Contest, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.ContestId, option.OpEqual, id)
	contest, err := dao.ContestStore.SelectOne(options)
	if err != nil {
		return Contest{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewContest().fromEntity(contest), &errors.NoError
}

func (_Query) SelectByUserId(userId uint64) ([]Contest, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.ContestUserId, option.OpEqual, userId)
	contests, err := dao.ContestStore.Select(options)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var result []Contest
	for _, contest := range contests {
		result = append(result, *NewContest().fromEntity(contest))
	}
	return result, &errors.NoError
}

func (_Query) SelectByStatus(status entity.ContestStatus) ([]Contest, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.ContestStatus, option.OpEqual, status)
	contests, err := dao.ContestStore.Select(options)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var result []Contest
	for _, contest := range contests {
		result = append(result, *NewContest().fromEntity(contest))
	}
	return result, &errors.NoError
}

func (_Query) Count(options *option.QueryOptions) (int64, error) {
	count, err := dao.ContestStore.Count(options)
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return count, &errors.NoError
}
