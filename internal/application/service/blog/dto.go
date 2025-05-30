package blog

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/blog"
	"STUOJ/internal/infrastructure/persistence/repository/dao"
	option "STUOJ/internal/infrastructure/persistence/repository/option"
	"STUOJ/internal/infrastructure/persistence/repository/querycontext"
	"STUOJ/pkg/utils"
	"time"
)

func params2Query(params request.QueryBlogParams) (query querycontext.BlogQueryContext) {
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
	if params.Problem != nil {
		ids, err := utils.StringToInt64Slice(*params.Problem)
		if err == nil {
			query.ProblemId.Set(ids)
		}
	}
	if params.User != nil {
		ids, err := utils.StringToInt64Slice(*params.User)
		if err == nil {
			query.UserId.Set(ids)
		}
	}
	if params.Status != nil {
		status, err := dao.StringToBlogStatusSlice(*params.Status)
		if err == nil {
			query.Status.Set(status)
		}
	}
	if params.Title != nil {
		query.Title.Set(*params.Title)
	}
	if params.Page != nil && params.Size != nil {
		query.Page = option.NewPagination(*params.Page, *params.Size)
	}
	if params.Order != nil && params.OrderBy != nil {
		query.Sort = option.NewSortQuery(*params.OrderBy, *params.Order)
	}
	return query
}

func domain2Resp(dm blog.Blog) (res response.BlogData) {
	res = response.BlogData{
		Id:         dm.Id.Value(),
		Title:      dm.Title.String(),
		Content:    dm.Content.String(),
		Status:     uint8(dm.Status.Value()),
		CreateTime: dm.CreateTime.String(),
		UpdateTime: dm.UpdateTime.String(),
	}
	return
}
