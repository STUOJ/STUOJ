package team

import (
	"STUOJ/internal/db/dao"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/domain/team/valueobject"
	"STUOJ/internal/errors"
	"fmt"
)

type Team struct {
	Id          uint64
	UserId      uint64
	ContestId   uint64
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
		return fmt.Errorf("用户ID不能为空")
	}
	if t.ContestId == 0 {
		return fmt.Errorf("比赛ID不能为空")
	}
	return nil
}

func (t *Team) toEntity() entity.Team {
	return entity.Team{
		Id:          t.Id,
		UserId:      t.UserId,
		ContestId:   t.ContestId,
		Name:        t.Name.String(),
		Description: t.Description.String(),
		Status:      t.Status,
	}
}

func (t *Team) fromEntity(team entity.Team) *Team {
	t.Id = team.Id
	t.UserId = team.UserId
	t.ContestId = team.ContestId
	t.Name = valueobject.NewName(team.Name)
	t.Description = valueobject.NewDescription(team.Description)
	t.Status = team.Status
	return t
}

func (t *Team) toOption() *option.QueryOptions {
	options := option.NewQueryOptions()
	options.Filters.Add(field.TeamId, option.OpEqual, t.Id)
	return options
}

func (t *Team) Create() (uint64, error) {
	if err := t.verify(); err != nil {
		return 0, errors.ErrValidation.WithMessage(err.Error())
	}
	team, err := dao.TeamStore.Insert(t.toEntity())
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return team.Id, &errors.NoError
}

func (t *Team) Update() error {
	var err error
	options := t.toOption()
	_, err = dao.TeamStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	if err := t.verify(); err != nil {
		return errors.ErrValidation.WithMessage(err.Error())
	}
	_, err = dao.TeamStore.Updates(t.toEntity(), options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return &errors.NoError
}

func (t *Team) Delete() error {
	options := t.toOption()
	_, err := dao.TeamStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	err = dao.TeamStore.Delete(options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return &errors.NoError
}

type Option func(*Team)

func NewTeam(options ...Option) *Team {
	team := &Team{}
	for _, option := range options {
		option(team)
	}
	return team
}

func WithId(id uint64) Option {
	return func(t *Team) {
		t.Id = id
	}
}

func WithUserId(userId uint64) Option {
	return func(t *Team) {
		t.UserId = userId
	}
}

func WithContestId(contestId uint64) Option {
	return func(t *Team) {
		t.ContestId = contestId
	}
}

func WithName(name string) Option {
	return func(t *Team) {
		t.Name = valueobject.NewName(name)
	}
}

func WithDescription(description string) Option {
	return func(t *Team) {
		t.Description = valueobject.NewDescription(description)
	}
}

func WithStatus(status entity.TeamStatus) Option {
	return func(t *Team) {
		t.Status = status
	}
}
