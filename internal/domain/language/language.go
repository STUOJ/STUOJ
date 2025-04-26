package language

//go:generate go run ../../../utils/gen/dto_gen.go language
//go:generate go run ../../../utils/gen/query_gen.go language

import (
	"STUOJ/internal/db/dao"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/domain/language/valueobject"
	"STUOJ/internal/errors"
	"STUOJ/internal/model/option"
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
	return int64(language.Id), &errors.NoError
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
	return &errors.NoError
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
	return &errors.NoError
}

type Option func(*Language)

func NewLanguage(option ...Option) *Language {
	l := &Language{}
	for _, opt := range option {
		opt(l)
	}
	return l
}

func WithId(id int64) Option {
	return func(l *Language) {
		l.Id = id
	}
}

func WithName(name string) Option {
	return func(l *Language) {
		l.Name = valueobject.NewName(name)
	}
}

func WithSerial(serial uint16) Option {
	return func(l *Language) {
		l.Serial = serial
	}
}

func WithMapId(mapId uint32) Option {
	return func(l *Language) {
		l.MapId = mapId
	}
}

func WithStatus(status entity.LanguageStatus) Option {
	return func(l *Language) {
		l.Status = status
	}
}
