package contest

//go:generate go run ../../../dev/gen/query_gen.go contest
//go:generate go run ../../../dev/gen/builder.go contest

import (
	dao2 "STUOJ/internal/infrastructure/repository/dao"
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/pkg/errors"
	"time"

	"STUOJ/internal/domain/contest/valueobject"
)

type Contest struct {
	Id          valueobject.Id
	UserId      valueobject.UserID
	Title       valueobject.Title
	Description valueobject.Description
	Status      valueobject.Status
	Format      valueobject.Format
	TeamSize    valueobject.TeamSize
	StartTime   valueobject.StartTime
	EndTime     valueobject.EndTime
	CreateTime  time.Time
	UpdateTime  time.Time
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
			ContestId: uint64(c.Id.Value()),
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
			ContestId: uint64(c.Id.Value()),
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
