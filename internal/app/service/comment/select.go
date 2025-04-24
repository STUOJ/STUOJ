package comment

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/app/dto/response"
	"STUOJ/internal/db/query"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/blog"
	"STUOJ/internal/domain/comment"
	"STUOJ/internal/domain/user"
	"STUOJ/internal/model"
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
	comments, _, err := comment.Query.Select(qc)
	if err != nil {
		return resp, err
	}

	userIds := make([]int64, len(comments))
	for i, c := range comments {
		userIds[i] = c.UserId
	}
	uqc := querycontext.UserQueryContext{}
	uqc.Id.Add(userIds...)
	uqc.Field = *query.UserSimpleField
	users, _, err := user.Query.Select(uqc)

	blogIds := make([]int64, len(comments))
	for i, c := range comments {
		blogIds[i] = c.BlogId
	}
	bqc := querycontext.BlogQueryContext{}
	bqc.Id.Add(blogIds...)
	bqc.Field = *query.BlogSimpleField
	blogs, _, err := blog.Query.Select(bqc)

	for _, u := range comments {
		respComment := domain2Resp(u)

		// 获取用户信息
		if u.UserId != 0 {
			respComment.User = response.Domain2UserSimpleData(users[u.UserId])
		}

		// 获取博客信息
		if u.BlogId != 0 {
			respComment.Blog = response.Domain2BlogSimpleData(blogs[u.BlogId])
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
