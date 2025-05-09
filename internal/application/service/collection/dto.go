package collection

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/collection"
	"STUOJ/internal/infrastructure/repository/dao"
	"STUOJ/internal/infrastructure/repository/querycontext"
	"STUOJ/internal/model/option"
	"STUOJ/pkg/utils"
	"time"
)

func params2Model(params request.QueryCollectionParams) (query querycontext.CollectionQueryContext) {
	if params.EndTime != nil {
		t, err := time.Parse("2006-01-02 15:04:05", *params.EndTime)
		if err == nil {
			query.EndTime.Set(t)
		}
	}
	if params.StartTime != nil {
		t, err := time.Parse("2006-01-02 15:04:05", *params.StartTime)
		if err == nil {
			query.StartTime.Set(t)
		}
	}
	if params.Title != nil {
		query.Title.Set(*params.Title)
	}
	if params.User != nil {
		ids, err := utils.StringToInt64Slice(*params.User)
		if err == nil {
			query.UserId.Set(ids)
		}
	}
	if params.Status != nil {
		status, err := dao.StringToCollectionStatusSlice(*params.Status)
		if err == nil {
			query.Status.Set(status)
		}
	}
	if params.Page != nil && params.Size != nil {
		query.Page = option.NewPagination(int64(*params.Page), int64(*params.Size))
	}
	if params.Order != nil && params.OrderBy != nil {
		query.Sort = option.NewSortQuery(*params.OrderBy, *params.Order)
	}

	if params.Problem != nil {
		problemIds, err := utils.StringToInt64Slice(*params.Problem)
		if err == nil {
			collection.WhereProblem(problemIds)(&query)
		}
	}

	return
}

func domain2response(c collection.Collection) response.CollectionData {
	return response.CollectionData{
		Id:          int64(c.Id.Value()),
		Title:       c.Title.String(),
		Description: c.Description.String(),
		Status:      int64(c.Status.Value()),
		CreateTime:  c.CreateTime.String(),
		UpdateTime:  c.UpdateTime.String(),
	}
}

func domain2listItemResponse(c collection.Collection) response.CollectionListItem {
	return response.CollectionListItem{
		Id:         int64(c.Id.Value()),
		Title:      c.Title.String(),
		Status:     int64(c.Status.Value()),
		CreateTime: c.CreateTime.String(),
		UpdateTime: c.UpdateTime.String(),
	}
}
