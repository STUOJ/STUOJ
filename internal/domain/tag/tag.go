package tag

//go:generate go run ../../../utils/gen/dto_gen.go tag
//go:generate go run ../../../utils/gen/query_gen.go tag

import (
	"STUOJ/internal/db/dao"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/domain/tag/valueobject"
	"STUOJ/internal/errors"
)

type Tag struct {
	Id   uint64
	Name valueobject.Name
}

func (t *Tag) verify() error {
	if err := t.Name.Verify(); err != nil {
		return err
	}
	return nil
}

func (t *Tag) toEntity() entity.Tag {
	return entity.Tag{
		Id:   t.Id,
		Name: t.Name.String(),
	}
}

func (t *Tag) fromEntity(tag entity.Tag) *Tag {
	t.Id = tag.Id
	t.Name = valueobject.NewName(tag.Name)
	return t
}

func (t *Tag) toOption() *option.QueryOptions {
	options := option.NewQueryOptions()
	options.Filters.Add(field.TagId, option.OpEqual, t.Id)
	return options
}

func (t *Tag) Create() (uint64, error) {
	if err := t.verify(); err != nil {
		return 0, errors.ErrValidation.WithMessage(err.Error())
	}
	tag, err := dao.TagStore.Insert(t.toEntity())
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return tag.Id, &errors.NoError
}

func (t *Tag) Update() error {
	options := t.toOption()
	_, err := dao.TagStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	if err := t.verify(); err != nil {
		return errors.ErrValidation.WithMessage(err.Error())
	}
	_, err = dao.TagStore.Updates(t.toEntity(), options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return &errors.NoError
}

func (t *Tag) Delete() error {
	options := t.toOption()
	_, err := dao.TagStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	err = dao.TagStore.Delete(options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return &errors.NoError
}

type Option func(*Tag)

func NewTag(options ...Option) *Tag {
	tag := &Tag{}
	for _, option := range options {
		option(tag)
	}
	return tag
}

func WithId(id uint64) Option {
	return func(t *Tag) {
		t.Id = id
	}
}

func WithName(name string) Option {
	return func(t *Tag) {
		t.Name = valueobject.NewName(name)
	}
}
