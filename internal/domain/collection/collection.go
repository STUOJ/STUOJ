package collection

//go:generate go run ../../../dev/gen/query_gen.go collection
//go:generate go run ../../../dev/gen/domain.go collection

import (
	dao2 "STUOJ/internal/infrastructure/repository/dao"
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/pkg/errors"
	"time"

	"STUOJ/internal/domain/collection/valueobject"
)

type Collection struct {
	Id          valueobject.Id
	UserId      valueobject.UserId
	Title       valueobject.Title
	Description valueobject.Description
	Status      valueobject.Status
	CreateTime  time.Time
	UpdateTime  time.Time
}

func (c *Collection) UpdateUser(userIds []int64) error {
	var err error
	options := c.toOption()
	_, err = dao2.CollectionStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	err = dao2.CollectionUserStore.Delete(options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	var errs []error
	for _, id := range userIds {
		_, err = dao2.CollectionUserStore.Insert(entity.CollectionUser{
			CollectionId: uint64(c.Id.Value()),
			UserId:       uint64(id),
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

func (c *Collection) UpdateProblem(problemIds []int64) error {
	var err error
	options := c.toOption()
	_, err = dao2.CollectionStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	err = dao2.CollectionProblemStore.Delete(options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	var errs []error
	var serial uint16 = 1
	for _, id := range problemIds {
		_, err = dao2.CollectionProblemStore.Insert(entity.CollectionProblem{
			CollectionId: uint64(c.Id.Value()),
			ProblemId:    uint64(id),
			Serial:       serial,
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
