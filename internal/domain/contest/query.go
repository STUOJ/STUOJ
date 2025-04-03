package contest

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

func (_Query) Select(model querymodel.ContestQueryModel) ([]Contest, error) {
	queryOptions := model.GenerateOptions()
	contests, err := dao.ContestStore.Select(queryOptions)
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
	queryOptions := option.NewQueryOptions()
	queryOptions.Filters.Add(field.ContestId, option.OpEqual, id)
	queryOptions.Field = query.ContestAllField
	contest, err := dao.ContestStore.SelectOne(queryOptions)
	if err != nil {
		return Contest{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewContest().fromEntity(contest), &errors.NoError
}

func (_Query) SelectSimpleById(id uint64) (Contest, error) {
	queryOptions := option.NewQueryOptions()
	queryOptions.Filters.Add(field.ContestId, option.OpEqual, id)
	queryOptions.Field = query.ContestSimpleField
	contest, err := dao.ContestStore.SelectOne(queryOptions)
	if err != nil {
		return Contest{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewContest().fromEntity(contest), &errors.NoError
}

func (_Query) Count(model querymodel.ContestQueryModel) (int64, error) {
	queryOptions := model.GenerateOptions()
	count, err := dao.ContestStore.Count(queryOptions)
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return count, &errors.NoError
}
