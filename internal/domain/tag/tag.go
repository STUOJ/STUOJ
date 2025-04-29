package tag

//go:generate go run ../../../dev/gen/dto_gen.go tag
//go:generate go run ../../../dev/gen/query_gen.go tag

import (
	"STUOJ/internal/domain/tag/valueobject"
	"STUOJ/internal/infrastructure/repository/dao"
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/infrastructure/repository/entity/field"
	"STUOJ/internal/model/option"
	"STUOJ/pkg/errors"
)

type Tag struct {
	Id   int64
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
		Id:   uint64(t.Id),
		Name: t.Name.String(),
	}
}

func (t *Tag) fromEntity(tag entity.Tag) *Tag {
	t.Id = int64(tag.Id)
	t.Name = valueobject.NewName(tag.Name)
	return t
}

func (t *Tag) toOption() *option.QueryOptions {
	options := option.NewQueryOptions()
	options.Filters.Add(field.TagId, option.OpEqual, t.Id)
	return options
}

func (t *Tag) Create() (int64, error) {
	if err := t.verify(); err != nil {
		return 0, errors.ErrValidation.WithMessage(err.Error())
	}
	tag, err := dao.TagStore.Insert(t.toEntity())
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return int64(tag.Id), &errors.NoError
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
