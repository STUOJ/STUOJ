package blog

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

func (*_Query) Select(model querymodel.BlogQueryModel) ([]Blog, error) {
	queryOptions := model.GenerateOptions()
	entityBlogs, err := dao.BlogStore.Select(queryOptions)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var blogs []Blog
	for _, entityBlog := range entityBlogs {
		blog := NewBlog().fromEntity(entityBlog)
		blogs = append(blogs, *blog)
	}
	return blogs, &errors.NoError
}

func (*_Query) SelectById(id uint64) (Blog, error) {
	queryOptions := option.NewQueryOptions()
	queryOptions.Filters.Add(field.BlogId, option.OpEqual, id)
	queryOptions.Field = query.BlogAllField
	entityBlog, err := dao.BlogStore.SelectOne(queryOptions)
	if err != nil {
		return Blog{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewBlog().fromEntity(entityBlog), &errors.NoError
}

func (*_Query) SelectSimpleById(id uint64) (Blog, error) {
	queryOptions := option.NewQueryOptions()
	queryOptions.Filters.Add(field.BlogId, option.OpEqual, id)
	queryOptions.Field = query.BlogSimpleField
	entityBlog, err := dao.BlogStore.SelectOne(queryOptions)
	if err != nil {
		return Blog{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewBlog().fromEntity(entityBlog), &errors.NoError
}

func (*_Query) Count(model querymodel.BlogQueryModel) (int64, error) {
	queryOptions := model.GenerateOptions()
	count, err := dao.BlogStore.Count(queryOptions)
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return count, &errors.NoError
}
