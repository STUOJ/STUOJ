package contest

//go:generate go run ../../../dev/gen/query_gen.go contest
//go:generate go run ../../../dev/gen/domain.go contest

import (
	"STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/internal/infrastructure/persistence/repository/dao"
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
