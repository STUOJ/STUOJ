package collection

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

func (*_Query) Select(model querymodel.CollectionQueryModel) ([]Collection, error) {
	queryOptions := model.GenerateOptions()
	queryOptions.Field = query.CollectionListItemField
	entityCollections, err := dao.CollectionStore.Select(queryOptions)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var collections []Collection
	for _, entityCollection := range entityCollections {
		collection := NewCollection().fromEntity(entityCollection)
		collections = append(collections, *collection)
	}
	return collections, &errors.NoError
}

func (*_Query) SelectById(id uint64) (Collection, error) {
	queryOptions := option.NewQueryOptions()
	queryOptions.Filters.Add(field.CollectionId, option.OpEqual, id)
	queryOptions.Field = query.CollectionAllField
	entityCollection, err := dao.CollectionStore.SelectOne(queryOptions)
	if err != nil {
		return Collection{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewCollection().fromEntity(entityCollection), &errors.NoError
}

func (*_Query) SelectSimpleById(id uint64) (Collection, error) {
	queryOptions := option.NewQueryOptions()
	queryOptions.Filters.Add(field.CollectionId, option.OpEqual, id)
	queryOptions.Field = query.CollectionSimpleField
	entityCollection, err := dao.CollectionStore.SelectOne(queryOptions)
	if err != nil {
		return Collection{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewCollection().fromEntity(entityCollection), &errors.NoError
}

func (*_Query) Count(model querymodel.CollectionQueryModel) (int64, error) {
	queryOptions := model.GenerateOptions()
	count, err := dao.CollectionStore.Count(queryOptions)
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return count, &errors.NoError
}

func (*_Query) SelectUserIds(id uint64) ([]uint64, error) {
	queryOptions := option.NewQueryOptions()
	queryOptions.Filters.Add(field.CollectionUserId, option.OpEqual, id)
	entityCollectionUsers, err := dao.CollectionUserStore.Select(queryOptions)
	if err != nil {
		return nil, errors.ErrNotFound.WithMessage(err.Error())
	}
	var userIds []uint64
	for _, entityCollectionUser := range entityCollectionUsers {
		userIds = append(userIds, entityCollectionUser.UserId)
	}
	return userIds, &errors.NoError
}

func (*_Query) SelectProblemIds(id uint64) ([]uint64, error) {
	queryOptions := option.NewQueryOptions()
	queryOptions.Filters.Add(field.CollectionProblemCollectionId, option.OpEqual, id)
	entityCollectionProblems, err := dao.CollectionProblemStore.Select(queryOptions)
	if err != nil {
		return nil, errors.ErrNotFound.WithMessage(err.Error())
	}
	var problemIds []uint64
	for _, entityCollectionProblem := range entityCollectionProblems {
		problemIds = append(problemIds, entityCollectionProblem.ProblemId)
	}
	return problemIds, &errors.NoError
}
