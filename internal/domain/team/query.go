package team

import (
	"STUOJ/internal/db/dao"
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/errors"
)

type _Query struct{}

var Query = new(_Query)

func (query *_Query) SelectById(id uint64) (*Team, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.TeamId, option.OpEqual, id)
	team, err := dao.TeamStore.SelectOne(options)
	if err != nil {
		return nil, errors.ErrNotFound.WithMessage(err.Error())
	}
	return NewTeam().fromEntity(team), nil
}

func (query *_Query) SelectByUserId(userId uint64) ([]*Team, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.TeamUserId, option.OpEqual, userId)
	teams, err := dao.TeamStore.Select(options)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var result []*Team
	for _, team := range teams {
		result = append(result, NewTeam().fromEntity(team))
	}
	return result, nil
}

func (query *_Query) SelectByContestId(contestId uint64) ([]*Team, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.TeamContestId, option.OpEqual, contestId)
	teams, err := dao.TeamStore.Select(options)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var result []*Team
	for _, team := range teams {
		result = append(result, NewTeam().fromEntity(team))
	}
	return result, nil
}

func (query *_Query) Select(options *option.QueryOptions) ([]*Team, error) {
	teams, err := dao.TeamStore.Select(options)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var result []*Team
	for _, team := range teams {
		result = append(result, NewTeam().fromEntity(team))
	}
	return result, nil
}

func (query *_Query) Count(options *option.QueryOptions) (int64, error) {
	count, err := dao.TeamStore.Count(options)
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return count, nil
}
