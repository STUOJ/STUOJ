package comment

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/blog"
	"STUOJ/internal/domain/comment"
	"STUOJ/internal/domain/user"
	entity "STUOJ/internal/infrastructure/repository/entity"
	query "STUOJ/internal/infrastructure/repository/query"
	querycontext "STUOJ/internal/infrastructure/repository/querycontext"
	model "STUOJ/internal/model"
	"slices"
)

type CommentPage struct {
	Comments []response.CommentData `json:"comments"`
	model.Page
}

// Select 查询所有评论
func Select(params request.QueryCommentParams, reqUser model.ReqUser) (CommentPage, error) {
	var resp CommentPage

	// 查询
	qc := params2Query(params)
	qc.Field.SelectId().SelectUserId().SelectBlogId().SelectStatus().SelectCreateTime().SelectUpdateTime()
	if !qc.Status.Exist() {
		qc.Status.Add(entity.CommentPublic)
	} else if slices.Contains(qc.Status.Value(), entity.CommentDeleted) && reqUser.Role < entity.RoleAdmin {
		qc.UserId.Set([]int64{reqUser.Id})
	}
	comments, _, err := comment.Query.Select(qc)
	if err != nil {
		return resp, err
	}

	userIds := make([]int64, len(comments))
	for i, c := range comments {
		userIds[i] = c.UserId.Value()
	}
	uqc := querycontext.UserQueryContext{}
	uqc.Id.Add(userIds...)
	uqc.Field = *query.UserSimpleField
	users, _, err := user.Query.Select(uqc)

	blogIds := make([]int64, len(comments))
	for i, c := range comments {
		blogIds[i] = c.BlogId.Value()
	}
	bqc := querycontext.BlogQueryContext{}
	bqc.Id.Add(blogIds...)
	bqc.Field = *query.BlogSimpleField
	blogs, _, err := blog.Query.Select(bqc)

	for _, u := range comments {
		respComment := domain2Resp(u)

		// 获取用户信息
		if u.UserId.Value() != 0 {
			respComment.User = response.Domain2UserSimpleData(users[u.UserId.Value()])
		}

		// 获取博客信息
		if u.BlogId.Value() != 0 {
			respComment.Blog = response.Domain2BlogSimpleData(blogs[u.BlogId.Value()])
		}

		resp.Comments = append(resp.Comments, respComment)
	}

	resp.Page.Page = qc.Page.Page
	resp.Size = qc.Page.PageSize
	resp.Page.Total, err = Count(params)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func Statistics(params request.CommentStatisticsParams, reqUser model.ReqUser) (response.StatisticsRes, error) {
	qc := params2Query(params.QueryCommentParams)
	qc.GroupBy = params.GroupBy
	resp, err := comment.Query.GroupCount(qc)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
