package blog

//go:generate go run ../../../dev/gen/query_gen.go blog
//go:generate go run ../../../dev/gen/builder.go blog

import (
	"STUOJ/internal/infrastructure/repository/dao"
	"STUOJ/internal/infrastructure/repository/entity/field"
	"STUOJ/internal/model/option"
	"STUOJ/pkg/errors"
	"time"

	"STUOJ/internal/domain/blog/valueobject"
)

type Blog struct {
	Id         valueobject.Id
	UserId     valueobject.UserId
	ProblemId  valueobject.ProblemId
	Title      valueobject.Title
	Content    valueobject.Content
	Status     valueobject.Status
	CreateTime time.Time
	UpdateTime time.Time
}

func (b *Blog) toOption() *option.QueryOptions {
	options := option.NewQueryOptions()
	options.Filters.Add(field.BlogId, option.OpEqual, b.Id)
	options.Field = b.existField()
	return options
}

func (b *Blog) Create() (int64, error) {
	b.CreateTime = time.Now()
	b.UpdateTime = time.Now()
	if err := b.verify(); err != nil {
		return 0, errors.ErrValidation.WithMessage(err.Error())
	}
	blog, err := dao.BlogStore.Insert(b.toEntity())
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return int64(blog.Id), nil
}

func (b *Blog) Update() error {
	var err error
	options := b.toOption()
	_, err = dao.BlogStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	b.UpdateTime = time.Now()
	if err := b.verify(); err != nil {
		return errors.ErrValidation.WithMessage(err.Error())
	}
	_, err = dao.BlogStore.Updates(b.toEntity(), options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return nil
}

func (b *Blog) Delete() error {
	options := b.toOption()
	_, err := dao.BlogStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	err = dao.BlogStore.Delete(options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return nil
}
