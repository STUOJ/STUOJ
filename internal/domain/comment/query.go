package comment

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

func (*_Query) Select(model querymodel.CommentQueryModel) ([]Comment, error) {
	queryOptions := model.GenerateOptions()
	entityComments, err := dao.CommentStore.Select(queryOptions)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var comments []Comment
	for _, entityComment := range entityComments {
		comment := NewComment().fromEntity(entityComment)
		comments = append(comments, *comment)
	}
	return comments, &errors.NoError
}

func (*_Query) SelectById(id uint64) (Comment, error) {
	queryOptions := option.NewQueryOptions()
	queryOptions.Filters.Add(field.CommentId, option.OpEqual, id)
	queryOptions.Field = query.CommentAllField
	entityComment, err := dao.CommentStore.SelectOne(queryOptions)
	if err != nil {
		return Comment{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewComment().fromEntity(entityComment), &errors.NoError
}

func (*_Query) SelectSimpleById(id uint64) (Comment, error) {
	queryOptions := option.NewQueryOptions()
	queryOptions.Filters.Add(field.CommentId, option.OpEqual, id)
	queryOptions.Field = query.CommentSimpleField
	entityComment, err := dao.CommentStore.SelectOne(queryOptions)
	if err != nil {
		return Comment{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewComment().fromEntity(entityComment), &errors.NoError
}

func (*_Query) SelectByUserId(userId uint64) ([]Comment, error) {
	queryOptions := option.NewQueryOptions()
	queryOptions.Filters.Add(field.CommentUserId, option.OpEqual, userId)
	entityComments, err := dao.CommentStore.Select(queryOptions)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var comments []Comment
	for _, entityComment := range entityComments {
		comment := NewComment().fromEntity(entityComment)
		comments = append(comments, *comment)
	}
	return comments, &errors.NoError
}

func (*_Query) SelectByBlogId(blogId uint64) ([]Comment, error) {
	queryOptions := option.NewQueryOptions()
	queryOptions.Filters.Add(field.CommentBlogId, option.OpEqual, blogId)
	entityComments, err := dao.CommentStore.Select(queryOptions)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var comments []Comment
	for _, entityComment := range entityComments {
		comment := NewComment().fromEntity(entityComment)
		comments = append(comments, *comment)
	}
	return comments, &errors.NoError
}

func (*_Query) Count(model querymodel.CommentQueryModel) (int64, error) {
	queryOptions := model.GenerateOptions()
	count, err := dao.CommentStore.Count(queryOptions)
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return count, &errors.NoError
}
