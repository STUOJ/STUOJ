package language

//go:generate go run ../../../dev/gen/query_gen.go language
//go:generate go run ../../../dev/gen/builder.go language

import (
	"STUOJ/internal/domain/language/valueobject"
	"STUOJ/internal/infrastructure/repository/dao"
	"STUOJ/pkg/errors"
)

type Language struct {
	Id     valueobject.Id
	Name   valueobject.Name
	Serial valueobject.Serial
	MapId  valueobject.MapId
	Status valueobject.Status
}

func (l *Language) Create() (int64, error) {
	if err := l.verify(); err != nil {
		return 0, errors.ErrValidation.WithMessage(err.Error())
	}
	language, err := dao.LanguageStore.Insert(l.toEntity())
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return int64(language.Id), nil
}

func (l *Language) Update() error {
	var err error
	options := l.toOption()
	_, err = dao.LanguageStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	if err := l.verify(); err != nil {
		return errors.ErrValidation.WithMessage(err.Error())
	}
	_, err = dao.LanguageStore.Updates(l.toEntity(), options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return nil
}

func (l *Language) Delete() error {
	options := l.toOption()
	_, err := dao.LanguageStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	err = dao.LanguageStore.Delete(options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return nil
}
