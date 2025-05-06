package team

//go:generate go run ../../../dev/gen/query_gen.go team
//go:generate go run ../../../dev/gen/builder.go team

import (
	"STUOJ/internal/domain/team/valueobject"
	dao2 "STUOJ/internal/infrastructure/repository/dao"
	"STUOJ/internal/infrastructure/repository/entity"
	field2 "STUOJ/internal/infrastructure/repository/entity/field"
	"STUOJ/internal/model/option"
	"STUOJ/pkg/errors"
)

type Team struct {
	Id          valueobject.Id
	UserId      valueobject.UserId
	ContestId   valueobject.ContestId
	Name        valueobject.Name
	Description valueobject.Description
	Status      valueobject.Status
}

func (t *Team) Create() (int64, error) {
	if err := t.verify(); err != nil {
		return 0, errors.ErrValidation.WithMessage(err.Error())
	}
	team, err := dao2.TeamStore.Insert(t.toEntity())
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return int64(team.Id), nil
}

func (t *Team) Update() error {
	var err error
	options := t.toOption()
	_, err = dao2.TeamStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	if err := t.verify(); err != nil {
		return errors.ErrValidation.WithMessage(err.Error())
	}
	_, err = dao2.TeamStore.Updates(t.toEntity(), options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return nil
}

func (t *Team) Delete() error {
	options := t.toOption()
	_, err := dao2.TeamStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	deleteTeamUserOptions := option.NewQueryOptions()
	deleteTeamUserOptions.Filters.Add(field2.TeamId, option.OpEqual, t.Id)
	err = dao2.TeamUserStore.Delete(deleteTeamUserOptions)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	err = dao2.TeamStore.Delete(options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}

	return nil
}

func (t *Team) JoinTeam(userId int64) error {
	options := t.toOption()
	_, err := dao2.TeamStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	_, err = dao2.TeamUserStore.Insert(entity.TeamUser{
		TeamId: uint64(t.Id.Value()),
		UserId: uint64(userId),
	})
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return nil
}

func (t *Team) QuitTeam(userId int64) error {
	options := t.toOption()
	_, err := dao2.TeamStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	deleteOptions := option.NewQueryOptions()
	deleteOptions.Filters.Add(field2.TeamId, option.OpEqual, t.Id)
	deleteOptions.Filters.Add(field2.UserId, option.OpEqual, userId)
	err = dao2.TeamStore.Delete(deleteOptions)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return nil
}
