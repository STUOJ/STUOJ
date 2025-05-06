package collection

//go:generate go run ../../../dev/gen/query_gen.go collection
//go:generate go run ../../../dev/gen/builder.go collection

import (
	dao2 "STUOJ/internal/infrastructure/repository/dao"
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/infrastructure/repository/entity/field"
	"STUOJ/internal/model/option"
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

func (c *Collection) toOption() *option.QueryOptions {
	options := option.NewQueryOptions()
	options.Filters.Add(field.CollectionId, option.OpEqual, c.Id)
	return options
}

func (c *Collection) Create() (int64, error) {
	c.CreateTime = time.Now()
	c.UpdateTime = time.Now()
	if err := c.verify(); err != nil {
		return 0, errors.ErrValidation.WithMessage(err.Error())
	}
	collection, err := dao2.CollectionStore.Insert(c.toEntity())
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return int64(collection.Id), nil
}

func (c *Collection) Update() error {
	var err error
	options := c.toOption()
	_, err = dao2.CollectionStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	c.UpdateTime = time.Now()
	if err := c.verify(); err != nil {
		return errors.ErrValidation.WithMessage(err.Error())
	}
	_, err = dao2.CollectionStore.Updates(c.toEntity(), options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return nil
}

func (c *Collection) Delete() error {
	options := c.toOption()
	_, err := dao2.CollectionStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	err = dao2.CollectionStore.Delete(options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return nil
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
