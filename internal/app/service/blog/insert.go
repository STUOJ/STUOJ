package blog

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/blog"
	"STUOJ/internal/domain/problem"
	"STUOJ/internal/errors"

	"STUOJ/internal/model"
)

// 插入博客
func Insert(req request.CreateBlogReq, reqUser model.ReqUser) (uint64, error) {
	blog := blog.NewBlog(blog.WithContent(req.Content),
		blog.WithTitle(req.Title),
		blog.WithUserID(uint64(reqUser.ID)),
		blog.WithStatus(entity.BlogStatus(req.Status)),
		blog.WithProblemID(uint64(req.ProblemID)),
	)
	if (blog.Status == entity.BlogBanned || blog.Status == entity.BlogNotice) && reqUser.Role < entity.RoleAdmin {
		return 0, errors.ErrUnauthorized.WithMessage("没有权限创建被封禁或公告的博客")
	}
	if blog.ProblemID != 0 {
		problemQueryContext := querycontext.ProblemQueryContext{}
		problemQueryContext.Id.Add(int64(blog.ProblemID))
		problemQueryContext.Field.SelectStatus()
		problem, _, err := problem.Query.SelectOne(problemQueryContext)
		if err != nil {
			return 0, errors.ErrNotFound.WithMessage("找不到对应的题目")
		}
		if problem.Status < entity.ProblemPublic {
			return 0, errors.ErrUnauthorized.WithMessage("没有权限创建对应题目未公开的博客")
		}
	}
	return blog.Create()
}
