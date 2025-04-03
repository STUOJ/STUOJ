package tag

import (
	"STUOJ/internal/db/dao"
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/errors"
	"STUOJ/internal/model/querymodel"
)

type _Query struct{}

var Query = new(_Query)

func (*_Query) Select(model querymodel.TagQueryModel) ([]Tag, error) {
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

func (*_Query) SelectById(id uint64) (Tag, error) {
	queryOptions := option.NewQueryOptions()
	queryOptions.Filters.Add(field.TagId, option.OpEqual, id)
	queryOptions.Field = query.TagAllField
	entityTag, err := dao.TagStore.SelectOne(queryOptions)
	if err != nil {
		return Tag{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewTag().fromEntity(entityTag), &errors.NoError
}

func (*_Query) Count(model querymodel.TagQueryModel) (int64, error) {
	queryOptions := model.GenerateOptions()
	count, err := dao.TagStore.Count(queryOptions)
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return count, &errors.NoError
}
