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
