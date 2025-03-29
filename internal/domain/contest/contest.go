package contest

import (
	"fmt"
	"time"

	"STUOJ/internal/db/dao"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/domain/contest/valueobject"
	"STUOJ/internal/errors"
)

type Contest struct {
	Id           uint64
	UserId       uint64
	CollectionId uint64
	Title        valueobject.Title
	Description  valueobject.Description
	Status       entity.ContestStatus
	Format       entity.ContestFormat
	TeamSize     uint8
	StartTime    time.Time
	EndTime      time.Time
	CreateTime   time.Time
	UpdateTime   time.Time
}

func (c *Contest) verify() error {
	if err := c.Title.Verify(); err != nil {
		return err
	}
	if err := c.Description.Verify(); err != nil {
		return err
	}
	if c.StartTime.After(c.EndTime) {
		return fmt.Errorf("比赛开始时间不能晚于结束时间！")
	}
	return nil
}

func (c *Contest) toEntity() entity.Contest {
	return entity.Contest{
		Id:           c.Id,
		UserId:       c.UserId,
		CollectionId: c.CollectionId,
		Status:       c.Status,
		Format:       c.Format,
		TeamSize:     c.TeamSize,
		StartTime:    c.StartTime,
		EndTime:      c.EndTime,
		CreateTime:   c.CreateTime,
		UpdateTime:   c.UpdateTime,
	}
}

func (c *Contest) fromEntity(contest entity.Contest) *Contest {
	c.Id = contest.Id
	c.UserId = contest.UserId
	c.CollectionId = contest.CollectionId
	c.Status = contest.Status
	c.Format = contest.Format
	c.TeamSize = contest.TeamSize
	c.StartTime = contest.StartTime
	c.EndTime = contest.EndTime
	c.CreateTime = contest.CreateTime
	c.UpdateTime = contest.UpdateTime
	return c
}

func (c *Contest) toOption() *option.QueryOptions {
	options := option.NewQueryOptions()
	options.Filters.Add(field.ContestId, option.OpEqual, c.Id)
	return options
}

func (c *Contest) Create() (uint64, error) {
	c.CreateTime = time.Now()
	c.UpdateTime = time.Now()
	if err := c.verify(); err != nil {
		return 0, errors.ErrValidation.WithMessage(err.Error())
	}
	contest, err := dao.ContestStore.Insert(c.toEntity())
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return contest.Id, &errors.NoError
}

func (c *Contest) Update() error {
	var err error
	options := c.toOption()
	_, err = dao.ContestStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	c.UpdateTime = time.Now()
	if err := c.verify(); err != nil {
		return errors.ErrValidation.WithMessage(err.Error())
	}
	_, err = dao.ContestStore.Updates(c.toEntity(), options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return &errors.NoError
}

func (c *Contest) Delete() error {
	options := c.toOption()
	_, err := dao.ContestStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	err = dao.ContestStore.Delete(options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return &errors.NoError
}

type Option func(*Contest)

func NewContest(option ...Option) *Contest {
	c := &Contest{}
	for _, opt := range option {
		opt(c)
	}
	return c
}

func WithId(id uint64) Option {
	return func(c *Contest) {
		c.Id = id
	}
}

func WithUserId(userId uint64) Option {
	return func(c *Contest) {
		c.UserId = userId
	}
}

func WithCollectionId(collectionId uint64) Option {
	return func(c *Contest) {
		c.CollectionId = collectionId
	}
}

func WithTitle(title string) Option {
	return func(c *Contest) {
		c.Title = valueobject.NewTitle(title)
	}
}

func WithDescription(description string) Option {
	return func(c *Contest) {
		c.Description = valueobject.NewDescription(description)
	}
}

func WithStatus(status entity.ContestStatus) Option {
	return func(c *Contest) {
		c.Status = status
	}
}

func WithFormat(format entity.ContestFormat) Option {
	return func(c *Contest) {
		c.Format = format
	}
}

func WithTeamSize(teamSize uint8) Option {
	return func(c *Contest) {
		c.TeamSize = teamSize
	}
}

func WithStartTime(startTime time.Time) Option {
	return func(c *Contest) {
		c.StartTime = startTime
	}
}

func WithEndTime(endTime time.Time) Option {
	return func(c *Contest) {
		c.EndTime = endTime
	}
}
