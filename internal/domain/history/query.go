package history

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

func (*_Query) SelectById(id uint64) (History, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.HistoryId, option.OpEqual, id)
	options.Field = query.HistoryAllField
	history, err := dao.HistoryStore.SelectOne(options)
	if err != nil {
		return History{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewHistory().fromEntity(history), &errors.NoError
}

func (*_Query) SelectSimpleById(id uint64) (History, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.HistoryId, option.OpEqual, id)
	options.Field = query.HistorySimpleField
	history, err := dao.HistoryStore.SelectOne(options)
	if err != nil {
		return History{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewHistory().fromEntity(history), &errors.NoError
}

func (*_Query) SelectByUserId(userId uint64) ([]History, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.HistoryUserId, option.OpEqual, userId)
	options.Field = query.HistorySimpleField
	histories, err := dao.HistoryStore.Select(options)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var result []History
	for _, history := range histories {
		result = append(result, *NewHistory().fromEntity(history))
	}
	return result, &errors.NoError
}

func (*_Query) SelectByProblemId(problemId uint64) ([]History, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.HistoryProblemId, option.OpEqual, problemId)
	options.Field = query.HistorySimpleField
	histories, err := dao.HistoryStore.Select(options)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var result []History
	for _, history := range histories {
		result = append(result, *NewHistory().fromEntity(history))
	}
	return result, &errors.NoError
}

func (*_Query) Select(model querymodel.HistoryQueryModel) ([]History, error) {
	options := model.GenerateOptions()
	options.Field = query.HistorySimpleField
	histories, err := dao.HistoryStore.Select(options)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var result []History
	for _, history := range histories {
		result = append(result, *NewHistory().fromEntity(history))
	}
	return result, &errors.NoError
}

func (*_Query) Count(model querymodel.HistoryQueryModel) (int64, error) {
	options := model.GenerateOptions()
	count, err := dao.HistoryStore.Count(options)
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return count, &errors.NoError
}
