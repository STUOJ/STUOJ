package contest

//go:generate go run ../../../dev/gen/query_gen.go contest
//go:generate go run ../../../dev/gen/builder.go contest

import (
	dao2 "STUOJ/internal/infrastructure/repository/dao"
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/infrastructure/repository/entity/field"
	"STUOJ/internal/model/option"
	"STUOJ/pkg/errors"
	"fmt"
	"time"

	"STUOJ/internal/domain/contest/valueobject"
)

type Contest struct {
	Id          int64
	UserId      int64
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
		Id:         uint64(c.Id),
		UserId:     uint64(c.UserId),
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
	c.Id = int64(contest.Id)
	c.UserId = int64(contest.UserId)
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

func (c *Contest) Create() (int64, error) {
	c.CreateTime = time.Now()
	c.UpdateTime = time.Now()
	if err := c.verify(); err != nil {
		return 0, errors.ErrValidation.WithMessage(err.Error())
	}
	contest, err := dao2.ContestStore.Insert(c.toEntity())
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return int64(contest.Id), nil
}

func (c *Contest) Update() error {
	var err error
	options := c.toOption()
	_, err = dao2.ContestStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	c.UpdateTime = time.Now()
	if err := c.verify(); err != nil {
		return errors.ErrValidation.WithMessage(err.Error())
	}
	_, err = dao2.ContestStore.Updates(c.toEntity(), options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return nil
}

func (c *Contest) Delete() error {
	options := c.toOption()
	_, err := dao2.ContestStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	err = dao2.ContestStore.Delete(options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return nil
}

func (c *Contest) UpdateUser(userIds []int64) error {
	var err error
	options := c.toOption()
	_, err = dao2.ContestStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	err = dao2.ContestUserStore.Delete(options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	var errs []error
	for _, id := range userIds {
		_, err = dao2.ContestUserStore.Insert(entity.ContestUser{
			ContestId: uint64(c.Id),
			UserId:    uint64(id),
		})
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errors.ErrInternalServer.WithErrors(errs)
	}
	return nil
}

func (c *Contest) UpdateProblem(problemIds []int64) error {
	var err error
	options := c.toOption()
	_, err = dao2.ContestStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	err = dao2.ContestProblemStore.Delete(options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	var errs []error
	var serial uint16 = 1
	for _, id := range problemIds {
		_, err = dao2.ContestProblemStore.Insert(entity.ContestProblem{
			ContestId: uint64(c.Id),
			ProblemId: uint64(id),
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
	return nil
}
