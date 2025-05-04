package language

//go:generate go run ../../../dev/gen/query_gen.go language
//go:generate go run ../../../dev/gen/builder.go language

import (
	"STUOJ/internal/domain/language/valueobject"
	"STUOJ/internal/infrastructure/repository/dao"
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/infrastructure/repository/entity/field"
	"STUOJ/internal/model/option"
	"STUOJ/pkg/errors"
)

type Language struct {
	Id     int64
	Name   valueobject.Name
	Serial uint16
	MapId  uint32
	Status entity.LanguageStatus
}

func (l *Language) verify() error {
	if err := l.Name.Verify(); err != nil {
		return err
	}
	return nil
}

func (l *Language) toEntity() entity.Language {
	return entity.Language{
		Id:     uint64(l.Id),
		Name:   l.Name.String(),
		Serial: l.Serial,
		MapId:  l.MapId,
		Status: l.Status,
	}
}

func (l *Language) fromEntity(language entity.Language) *Language {
	l.Id = int64(language.Id)
	l.Name = valueobject.NewName(language.Name)
	l.Serial = language.Serial
	l.MapId = language.MapId
	l.Status = language.Status
	return l
}

func (l *Language) toOption() *option.QueryOptions {
	options := option.NewQueryOptions()
	options.Filters.Add(field.LanguageId, option.OpEqual, l.Id)
	return options
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
