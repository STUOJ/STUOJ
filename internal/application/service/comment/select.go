package comment

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/blog"
	"STUOJ/internal/domain/comment"
	"STUOJ/internal/domain/user"
	entity2 "STUOJ/internal/infrastructure/repository/entity"
	query2 "STUOJ/internal/infrastructure/repository/query"
	querycontext2 "STUOJ/internal/infrastructure/repository/querycontext"
	model2 "STUOJ/internal/model"
	"slices"
)

type CommentPage struct {
	Comments []response.CommentData `json:"comments"`
	model2.Page
}

// Select 查询所有评论
func Select(params request.QueryCommentParams, reqUser model2.ReqUser) (CommentPage, error) {
	var resp CommentPage

	// 查询
	qc := params2Query(params)
	qc.Field.SelectId().SelectUserId().SelectBlogId().SelectStatus().SelectCreateTime().SelectUpdateTime()
	if !qc.Status.Exist() {
		qc.Status.Add(int64(entity2.CommentPublic))
	} else if slices.Contains(qc.Status.Value(), int64(entity2.CommentDeleted)) && reqUser.Role < entity2.RoleAdmin {
		qc.UserId.Set([]int64{reqUser.Id})
	}
	comments, _, err := comment.Query.Select(qc)
	if err != nil {
		return resp, err
	}

	userIds := make([]int64, len(comments))
	for i, c := range comments {
		userIds[i] = c.UserId
	}
	uqc := querycontext2.UserQueryContext{}
	uqc.Id.Add(userIds...)
	uqc.Field = *query2.UserSimpleField
	users, _, err := user.Query.Select(uqc)

	blogIds := make([]int64, len(comments))
	for i, c := range comments {
		blogIds[i] = c.BlogId
	}
	bqc := querycontext2.BlogQueryContext{}
	bqc.Id.Add(blogIds...)
	bqc.Field = *query2.BlogSimpleField
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
