package contest

import (
	"STUOJ/internal/db/dao"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/errors"
)

func GetContestById(id uint64) (*Contest, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.ContestId, option.OpEqual, id)
	contest, err := dao.ContestStore.SelectOne(options)
	if err != nil {
		return nil, errors.ErrNotFound.WithMessage(err.Error())
	}
	return NewContest().fromEntity(contest), nil
}

func GetContestByUserId(userId uint64) ([]*Contest, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.ContestUserId, option.OpEqual, userId)
	contests, err := dao.ContestStore.Select(options)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var result []*Contest
	for _, contest := range contests {
		result = append(result, NewContest().fromEntity(contest))
	}
	return result, nil
}

func GetContestByStatus(status entity.ContestStatus) ([]*Contest, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.ContestStatus, option.OpEqual, status)
	contests, err := dao.ContestStore.Select(options)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var result []*Contest
	for _, contest := range contests {
		result = append(result, NewContest().fromEntity(contest))
	}
	return result, nil
}

func GetContestList(options *option.QueryOptions) ([]*Contest, error) {
	contests, err := dao.ContestStore.Select(options)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var result []*Contest
	for _, contest := range contests {
		result = append(result, NewContest().fromEntity(contest))
	}
	return result, nil
}

func GetContestCount(options *option.QueryOptions) (int64, error) {
	count, err := dao.ContestStore.Count(options)
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return count, nil
}
