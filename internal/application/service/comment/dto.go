package comment

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/comment"
	"STUOJ/internal/infrastructure/repository/querycontext"
	"STUOJ/internal/model/option"
	"STUOJ/pkg/utils"
	"time"
)

func domain2Resp(dm comment.Comment) (resp response.CommentData) {
	resp = response.CommentData{
		Id:         dm.Id.Value(),
		Content:    dm.Content.String(),
		Status:     int64(dm.Status.Value()),
		CreateTime: dm.CreateTime.String(),
		UpdateTime: dm.UpdateTime.String(),
	}
	return
}

func params2Query(params request.QueryCommentParams) (query querycontext.CommentQueryContext) {
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
	if params.Status != nil {
		ids, err := utils.StringToInt64Slice(*params.Status)
		if err == nil {
			query.Status.Set(ids)
		}
	}
	if params.User != nil {
		query.UserId.Add(*params.User)
	}
	if params.Blog != nil {
		query.BlogId.Set(*params.Blog)
	}
	if params.Page != nil && params.Size != nil {
		query.Page = option.NewPagination(*params.Page, *params.Size)
	}
	if params.Order != nil && params.OrderBy != nil {
		query.Sort = option.NewSortQuery(*params.Order, *params.OrderBy)
	}
	return query
}
