package tag

import (
	"STUOJ/internal/db/dao"
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/errors"
	"STUOJ/internal/model/querymodel"
)

type _Query struct{}

var Query = new(_Query)

func (query *_Query) Select(model querymodel.TagQueryModel) ([]Tag, error) {
	queryOptions := model.GenerateOptions()
	entityTags, err := dao.TagStore.Select(queryOptions)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var tags []Tag
	for _, entityTag := range entityTags {
		tag := NewTag().fromEntity(entityTag)
		tags = append(tags, *tag)
	}
	return tags, nil
}

func (query *_Query) SelectById(id uint64) (Tag, error) {
	queryOptions := option.NewQueryOptions()
	queryOptions.Filters.Add(field.TagId, option.OpEqual, id)
	entityTag, err := dao.TagStore.SelectOne(queryOptions)
	if err != nil {
		return Tag{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewTag().fromEntity(entityTag), nil
}

func (query *_Query) SelectByName(name string) (Tag, error) {
	queryOptions := option.NewQueryOptions()
	queryOptions.Filters.Add(field.TagName, option.OpEqual, name)
	entityTag, err := dao.TagStore.SelectOne(queryOptions)
	if err != nil {
		return Tag{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewTag().fromEntity(entityTag), nil
}

func (query *_Query) Count(model querymodel.TagQueryModel) (int64, error) {
	queryOptions := model.GenerateOptions()
	count, err := dao.TagStore.Count(queryOptions)
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return count, nil
}
