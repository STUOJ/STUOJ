package team

//go:generate go run ../../../dev/gen/dto_gen.go team
//go:generate go run ../../../dev/gen/query_gen.go team
//go:generate go run ../../../dev/gen/builder.go team

import (
	"STUOJ/internal/domain/team/valueobject"
	dao2 "STUOJ/internal/infrastructure/repository/dao"
	"STUOJ/internal/infrastructure/repository/entity"
	field2 "STUOJ/internal/infrastructure/repository/entity/field"
	"STUOJ/internal/model/option"
	"STUOJ/pkg/errors"
	"fmt"
)

type Team struct {
	Id          int64
	UserId      int64
	ContestId   int64
	Name        valueobject.Name
	Description valueobject.Description
	Status      entity.TeamStatus
}

func (t *Team) verify() error {
	if err := t.Name.Verify(); err != nil {
		return err
	}
	if err := t.Description.Verify(); err != nil {
		return err
	}
	if t.UserId == 0 {
		return fmt.Errorf("用户Id不能为空")
	}
	if t.ContestId == 0 {
		return fmt.Errorf("比赛Id不能为空")
	}
	return nil
}

func (t *Team) toEntity() entity.Team {
	return entity.Team{
		Id:          uint64(t.Id),
		UserId:      uint64(t.UserId),
		ContestId:   uint64(t.ContestId),
		Name:        t.Name.String(),
		Description: t.Description.String(),
		Status:      t.Status,
	}
}

func (t *Team) fromEntity(team entity.Team) *Team {
	t.Id = int64(team.Id)
	t.UserId = int64(team.UserId)
	t.ContestId = int64(team.ContestId)
	t.Name = valueobject.NewName(team.Name)
	t.Description = valueobject.NewDescription(team.Description)
	t.Status = team.Status
	return t
}

func (t *Team) toOption() *option.QueryOptions {
	options := option.NewQueryOptions()
	options.Filters.Add(field2.TeamId, option.OpEqual, t.Id)
	return options
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
		TeamId: uint64(t.Id),
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
