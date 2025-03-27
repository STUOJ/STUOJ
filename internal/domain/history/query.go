package history

import (
	"STUOJ/internal/db/dao"
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/errors"
)

type _Query struct{}

var Query = new(_Query)

func (query *_Query) SelectById(id uint64) (*History, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.HistoryId, option.OpEqual, id)
	history, err := dao.HistoryStore.SelectOne(options)
	if err != nil {
		return nil, errors.ErrNotFound.WithMessage(err.Error())
	}
	return NewHistory().fromEntity(history), nil
}

func (query *_Query) SelectByUserId(userId uint64) ([]*History, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.HistoryUserId, option.OpEqual, userId)
	histories, err := dao.HistoryStore.Select(options)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var result []*History
	for _, history := range histories {
		result = append(result, NewHistory().fromEntity(history))
	}
	return result, nil
}

func (query *_Query) SelectByProblemId(problemId uint64) ([]*History, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.HistoryProblemId, option.OpEqual, problemId)
	histories, err := dao.HistoryStore.Select(options)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var result []*History
	for _, history := range histories {
		result = append(result, NewHistory().fromEntity(history))
	}
	return result, nil
}

func (query *_Query) Select(options *option.QueryOptions) ([]*History, error) {
	histories, err := dao.HistoryStore.Select(options)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var result []*History
	for _, history := range histories {
		result = append(result, NewHistory().fromEntity(history))
	}
	return result, nil
}

func (query *_Query) Count(options *option.QueryOptions) (int64, error) {
	count, err := dao.HistoryStore.Count(options)
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return count, nil
}
