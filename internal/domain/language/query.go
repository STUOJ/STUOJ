package language

import (
	"STUOJ/internal/db/dao"
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/errors"
)

type _Query struct{}

var Query = new(_Query)

func (*_Query) SelectById(id uint64) (*Language, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.LanguageId, option.OpEqual, id)
	language, err := dao.LanguageStore.SelectOne(options)
	if err != nil {
		return nil, errors.ErrNotFound.WithMessage(err.Error())
	}
	return NewLanguage().fromEntity(language), &errors.NoError
}

func (*_Query) SelectByName(name string) (Language, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.LanguageName, option.OpEqual, name)
	language, err := dao.LanguageStore.SelectOne(options)
	if err != nil {
		return Language{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewLanguage().fromEntity(language), &errors.NoError
}

func (*_Query) Select(options *option.QueryOptions) ([]Language, error) {
	languages, err := dao.LanguageStore.Select(options)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var result []Language
	for _, language := range languages {
		result = append(result, *NewLanguage().fromEntity(language))
	}
	return result, &errors.NoError
}

func (*_Query) Count(options *option.QueryOptions) (int64, error) {
	count, err := dao.LanguageStore.Count(options)
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return count, &errors.NoError
}
