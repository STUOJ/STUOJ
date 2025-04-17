package blog

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/app/dto/response"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/blog"
	"STUOJ/utils"
	"time"
)

func params2Model(params request.QueryBlogParams) (query querycontext.BlogQueryContext) {
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
		if err != nil {
			query.ProblemId.Set(ids)
		}
	}
	if params.User != nil {
		ids, err := utils.StringToInt64Slice(*params.User)
		if err != nil {
			query.UserId.Set(ids)
		}
	}
	if params.Status != nil {
		ids, err := utils.StringToInt64Slice(*params.Status)
		if err != nil {
			query.Status.Set(ids)
		}
	}
	if params.Title != nil {
		query.Title.Set(*params.Title)
	}
	if params.Page != nil && params.Size != nil {
		query.Page = option.NewPagination(*params.Page, *params.Size)
	}
	if params.Order != nil && params.OrderBy != nil {
		query.Sort = option.NewSortQuery(*params.Order, *params.OrderBy)
	}
	return query
}

func domain2response(domainBlog blog.Blog) (res response.BlogData) {
	res = response.BlogData{
		ID:         int64(domainBlog.ID),
		Title:      domainBlog.Title.String(),
		Content:    domainBlog.Content.String(),
		Status:     int64(domainBlog.Status),
		CreateTime: domainBlog.CreateTime.String(),
		UpdateTime: domainBlog.UpdateTime.String(),
	}
	return
}
