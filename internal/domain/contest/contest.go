package contest

//go:generate go run ../../../utils/gen/dto_gen.go contest
//go:generate go run ../../../utils/gen/query_gen.go contest

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
	Id          uint64
	UserId      uint64
	Title       valueobject.Title
	Description valueobject.Description
	Status      entity.ContestStatus
	Format      entity.ContestFormat
	TeamSize    uint8
	StartTime   time.Time
	EndTime     time.Time
	CreateTime  time.Time
	UpdateTime  time.Time
}

func (c *Contest) verify() error {
	if c.UserId == 0 {
		return fmt.Errorf("用户Id不能为空")
	}
	if !entity.ContestStatus(c.Status).IsValid() {
		return fmt.Errorf("比赛状态不合法")
	}
	if !entity.ContestFormat(c.Format).IsValid() {
		return fmt.Errorf("比赛格式不合法")
	}
	if c.StartTime.IsZero() {
		return fmt.Errorf("比赛开始时间不能为空")
	}
	if c.EndTime.IsZero() {
		return fmt.Errorf("比赛结束时间不能为空")
	}
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
		Id:         c.Id,
		UserId:     c.UserId,
		Status:     c.Status,
		Format:     c.Format,
		TeamSize:   c.TeamSize,
		StartTime:  c.StartTime,
		EndTime:    c.EndTime,
		CreateTime: c.CreateTime,
		UpdateTime: c.UpdateTime,
	}
}

func (c *Contest) fromEntity(contest entity.Contest) *Contest {
	c.Id = contest.Id
	c.UserId = contest.UserId
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

func (c *Contest) UpdateUser(userIds []uint64) error {
	var err error
	options := c.toOption()
	_, err = dao.ContestStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	err = dao.ContestUserStore.Delete(options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	var errs []error
	for _, id := range userIds {
		_, err = dao.ContestUserStore.Insert(entity.ContestUser{
			ContestId: c.Id,
			UserId:    id,
		})
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errors.ErrInternalServer.WithErrors(errs)
	}
	return &errors.NoError
}

func (c *Contest) UpdateProblem(problemIds []uint64) error {
	var err error
	options := c.toOption()
	_, err = dao.ContestStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	err = dao.ContestProblemStore.Delete(options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	var errs []error
	var serial uint16 = 1
	for _, id := range problemIds {
		_, err = dao.ContestProblemStore.Insert(entity.ContestProblem{
			ContestId: c.Id,
			ProblemId: id,
			Serial:    serial,
		})
		if err != nil {
			errs = append(errs, err)
		}
		serial++
	}
	if len(errs) > 0 {
		return errors.ErrInternalServer.WithErrors(errs)
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
