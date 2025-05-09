package team

//go:generate go run ../../../dev/gen/query_gen.go team
//go:generate go run ../../../dev/gen/domain.go team

import (
	"STUOJ/internal/domain/team/valueobject"
	"STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/internal/infrastructure/persistence/entity/field"
	"STUOJ/internal/infrastructure/persistence/repository/dao"
	option2 "STUOJ/internal/infrastructure/persistence/repository/option"
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
	_, err := dao.TeamStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	_, err = dao.TeamUserStore.Insert(entity.TeamUser{
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
	_, err := dao.TeamStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	deleteOptions := option2.NewQueryOptions()
	deleteOptions.Filters.Add(field.TeamId, option2.OpEqual, t.Id)
	deleteOptions.Filters.Add(field.UserId, option2.OpEqual, userId)
	err = dao.TeamStore.Delete(deleteOptions)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return nil
}
