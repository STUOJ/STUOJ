package tag

//go:generate go run ../../../dev/gen/query_gen.go tag
//go:generate go run ../../../dev/gen/builder.go tag

import (
	"STUOJ/internal/domain/tag/valueobject"
	"STUOJ/internal/infrastructure/repository/dao"
	"STUOJ/pkg/errors"
)

type Tag struct {
	Id   valueobject.Id
	Name valueobject.Name
}

func (t *Tag) Create() (int64, error) {
	if err := t.verify(); err != nil {
		return 0, errors.ErrValidation.WithMessage(err.Error())
	}
	tag, err := dao.TagStore.Insert(t.toEntity())
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return int64(tag.Id), nil
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
	return nil
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
	return nil
}
