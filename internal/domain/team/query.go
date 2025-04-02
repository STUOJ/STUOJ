package team

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

func (*_Query) SelectById(id uint64) (Team, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.TeamId, option.OpEqual, id)
	team, err := dao.TeamStore.SelectOne(options)
	if err != nil {
		return Team{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewTeam().fromEntity(team), &errors.NoError
}

func (*_Query) SelectSimpleById(id uint64) (Team, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.TeamId, option.OpEqual, id)
	options.Field = query.TeamSimpleField
	team, err := dao.TeamStore.SelectOne(options)
	if err != nil {
		return Team{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewTeam().fromEntity(team), &errors.NoError
}

func (*_Query) SelectByUserId(userId uint64) ([]Team, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.TeamUserId, option.OpEqual, userId)
	options.Field = query.TeamAllField
	teams, err := dao.TeamStore.Select(options)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var result []Team
	for _, team := range teams {
		result = append(result, *NewTeam().fromEntity(team))
	}
	return result, &errors.NoError
}

func (*_Query) SelectByContestId(contestId uint64) ([]Team, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.TeamContestId, option.OpEqual, contestId)
	options.Field = query.TeamAllField
	teams, err := dao.TeamStore.Select(options)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var result []Team
	for _, team := range teams {
		result = append(result, *NewTeam().fromEntity(team))
	}
	return result, &errors.NoError
}

func (*_Query) Select(model querymodel.TeamQueryModel) ([]Team, error) {
	queryOptions := model.GenerateOptions()
	queryOptions.Field = query.TeamAllField
	teams, err := dao.TeamStore.Select(queryOptions)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var result []Team
	for _, team := range teams {
		result = append(result, *NewTeam().fromEntity(team))
	}
	return result, &errors.NoError
}

func (*_Query) Count(model querymodel.TeamQueryModel) (int64, error) {
	queryOptions := model.GenerateOptions()
	count, err := dao.TeamStore.Count(queryOptions)
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return count, &errors.NoError
}
